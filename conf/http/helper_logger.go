package http

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/saitofun/qkit/x/misc/timer"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/saitofun/qkit/kit/httptransport/httpx"
	"github.com/saitofun/qkit/kit/metax"
)

type LogRoundTripper struct {
	logger *logrus.Entry
	next   http.RoundTripper
}

func NewLogRoundTripper(logger *logrus.Entry) func(http.RoundTripper) http.RoundTripper {
	return func(rt http.RoundTripper) http.RoundTripper {
		return &LogRoundTripper{logger, rt}
	}
}

func (rt *LogRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	cost := timer.Start()

	ctx := req.Context()

	// inject b3 form context
	b3.New(b3.WithInjectEncoding(b3.B3SingleHeader)).
		Inject(ctx, propagation.HeaderCarrier(req.Header))

	rsp, err := rt.next.RoundTrip(req)

	level, _ := logrus.ParseLevel(strings.ToLower(req.Header.Get("x-log-level")))
	if level == logrus.PanicLevel {
		level = rt.logger.Logger.Level
	}

	if err == nil {
		// extract b3 to ctx
		b3.New().Extract(ctx, propagation.HeaderCarrier(rsp.Header))
	}

	logger := rt.logger.WithContext(ctx).WithFields(logrus.Fields{
		"cost":   fmt.Sprintf("%0.3fms", float64(cost()/time.Millisecond)),
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

func LogHandler(logger *logrus.Entry, tracer trace.Tracer) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := b3.New().Extract(
				req.Context(),
				propagation.HeaderCarrier(req.Header),
			)
			cost := timer.StartSpan()
			ctx, span := tracer.Start(
				ctx,
				"UnknownOperation",
				trace.WithTimestamp(cost.StartedAt()),
			)

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

			next.ServeHTTP(lrw, req.WithContext(ctx))

			operator := metax.ParseMeta(lrw.Header().Get("X-Meta")).Get("operator")
			if operator != "" {
				span.SetName(operator)
			}

			level, _ := logrus.ParseLevel(strings.ToLower(req.Header.Get("x-log-level")))
			if level == logrus.PanicLevel {
				level = logger.Logger.Level
			}

			entry := logger.WithContext(metax.ContextWithMeta(
				ctx,
				metax.ParseMeta(lrw.Header().Get("X-Meta")),
			))

			header := req.Header

			fields := logrus.Fields{
				"tag":         "access",
				"remote_ip":   httpx.ClientIP(req),
				"cost":        fmt.Sprintf("%0.3fms", float64(cost.Cost().Milliseconds())),
				"method":      req.Method,
				"request_uri": omitAuthorization(req.URL),
				"user_agent":  header.Get(httpx.HeaderUserAgent),
				"status":      lrw.statusCode,
			}

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
