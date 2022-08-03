package http

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/saitofun/qkit/kit/httptransport/httpx"
	"github.com/saitofun/qkit/kit/metax"
)

func NewLogRoundTripper(logger *logrus.Entry) func(http.RoundTripper) http.RoundTripper {
	return func(roundTripper http.RoundTripper) http.RoundTripper {
		return &LogRoundTripper{
			logger: logger,
			next:   roundTripper,
		}
	}
}

type LogRoundTripper struct {
	logger *logrus.Entry
	next   http.RoundTripper
}

func (rt *LogRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	startedAt := time.Now()

	ctx := req.Context()

	// inject h3 form context
	b3.New(b3.WithInjectEncoding(b3.B3SingleHeader)).
		Inject(ctx, propagation.HeaderCarrier(req.Header))

	rsp, err := rt.next.RoundTrip(req)

	level, _ := logrus.ParseLevel(strings.ToLower(req.Header.Get("x-log-level")))
	if level == logrus.PanicLevel {
		level = rt.logger.Logger.Level
	}

	cost := time.Since(startedAt)
	if err == nil {
		// extract b3 to ctx
		ctx = b3.New().Extract(ctx, propagation.HeaderCarrier(rsp.Header))
	}

	logger := rt.logger.WithContext(ctx).WithFields(logrus.Fields{
		"cost":   fmt.Sprintf("%0.3fms", float64(cost/time.Millisecond)),
		"method": req.Method,
		"url":    omitAuthorization(req.URL),
	})

	if err == nil {
		if level >= logrus.InfoLevel {
			logger.Infof("success")
		}
	} else {
		if level >= logrus.WarnLevel {
			logger.Warnf("do http request failed %s", err)
		}
	}

	return rsp, err
}

func LogHandler(logger *logrus.Entry, tracer trace.Tracer) func(handler http.Handler) http.Handler {
	return func(nextHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := req.Context()

			ctx = b3.New().Extract(ctx, propagation.HeaderCarrier(req.Header))
			startAt := time.Now()

			ctx, span := tracer.Start(ctx, "UnknownOperation", trace.WithTimestamp(startAt))
			defer func() {
				span.End(trace.WithTimestamp(time.Now()))
			}()

			lrw := newLoggerResponseWriter(rw)

			// for async pick
			b3.New(b3.WithInjectEncoding(b3.B3SingleHeader)).
				Inject(ctx, propagation.HeaderCarrier(lrw.Header()))

			ctx = metax.ContextWithMeta(
				ctx,
				metax.ParseMeta(span.SpanContext().TraceID().String()),
			)

			nextHandler.ServeHTTP(lrw, req.WithContext(ctx))

			op := metax.ParseMeta(lrw.Header().Get("X-Meta")).Get("operator")
			if op != "" {
				span.SetName(op)
			}

			level, _ := logrus.ParseLevel(strings.ToLower(req.Header.Get("x-log-level")))
			if level == logrus.PanicLevel {
				level = logger.Logger.Level
			}

			duration := time.Since(startAt)

			entry := logger.WithContext(metax.ContextWithMeta(
				ctx,
				metax.ParseMeta(lrw.Header().Get("X-Meta")),
			))

			header := req.Header

			fields := logrus.Fields{
				"tag":         "access",
				"remote_ip":   httpx.ClientIP(req),
				"cost":        fmt.Sprintf("%0.3fms", float64(duration/time.Millisecond)),
				"method":      req.Method,
				"request_uri": omitAuthorization(req.URL),
				"user_agent":  header.Get(httpx.HeaderUserAgent),
			}

			fields["status"] = lrw.statusCode

			if lrw.errMsg.Len() > 0 {
				if lrw.statusCode >= http.StatusInternalServerError {
					if level >= logrus.ErrorLevel {
						entry.WithFields(fields).Error(lrw.errMsg.String())
					}
				} else {
					if level >= logrus.WarnLevel {
						entry.WithFields(fields).Warn(lrw.errMsg.String())
					}
				}
			} else {
				if level >= logrus.InfoLevel {
					entry.WithFields(fields).Info()
				}
			}
		})
	}
}

func newLoggerResponseWriter(rw http.ResponseWriter) *loggerResponseWriter {
	h, hok := rw.(http.Hijacker)
	if !hok {
		h = nil
	}

	f, fok := rw.(http.Flusher)
	if !fok {
		f = nil
	}

	return &loggerResponseWriter{
		ResponseWriter: rw,
		Hijacker:       h,
		Flusher:        f,
	}
}

type loggerResponseWriter struct {
	http.ResponseWriter
	http.Hijacker
	http.Flusher

	headerWritten bool
	statusCode    int
	errMsg        bytes.Buffer
}

func (rw *loggerResponseWriter) Header() http.Header {
	return rw.ResponseWriter.Header()
}

func (rw *loggerResponseWriter) WriteHeader(statusCode int) {
	rw.writeHeader(statusCode)
}

func (rw *loggerResponseWriter) Write(data []byte) (int, error) {
	if rw.statusCode >= http.StatusBadRequest {
		rw.errMsg.Write(data)
	}
	return rw.ResponseWriter.Write(data)
}

func (rw *loggerResponseWriter) writeHeader(statusCode int) {
	if !rw.headerWritten {
		rw.ResponseWriter.WriteHeader(statusCode)
		rw.statusCode = statusCode
		rw.headerWritten = true
	}
}

func omitAuthorization(u *url.URL) string {
	query := u.Query()
	query.Del("authorization")
	u.RawQuery = query.Encode()
	return u.String()
}
