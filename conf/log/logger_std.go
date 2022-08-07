package log

import (
	"context"
	"fmt"
	"log"

	"github.com/saitofun/qkit/x/ptrx"
)

func Std() Logger { return &std{lvl: DebugLevel, lvlKey: "@lvl"} }

func StdContext(ctx context.Context) (context.Context, Logger) {
	l := Std()
	return WithLogger(ctx, l), l
}

type std struct {
	lvl    Level
	lvlKey string
	spans  []string
	kvs    []interface{}
}

func (l *std) SetLevel(lvl Level) Logger {
	logger := ptrx.Ptr(*l)
	logger.lvl = lvl
	return logger
}

func (l *std) SetLevelKey(k string) { l.lvlKey = k }

func (l *std) WithValues(kvs ...interface{}) Logger {
	return &std{
		lvl:   l.lvl,
		spans: l.spans,
		kvs:   append(l.kvs, kvs...),
	}
}

func (l *std) Start(ctx context.Context, name string, kvs ...interface{}) (context.Context, Logger) {
	return ctx, &std{
		lvl:   l.lvl,
		spans: append(l.spans, name),
		kvs:   append(l.kvs, kvs...),
	}
}

func (l *std) End() {
	if len(l.spans) != 0 {
		l.spans = l.spans[0 : len(l.spans)-1]
	}
}

func (l *std) Trace(format string, args ...interface{}) {
	if TraceLevel <= l.lvl {
		log.Println(append(KeyValues(append(l.kvs, l.lvlKey, "trace")),
			fmt.Sprintf(format, args...))...)
	}
}

func (l *std) Debug(format string, args ...interface{}) {
	if DebugLevel <= l.lvl {
		log.Println(append(KeyValues(append(l.kvs, l.lvlKey, "debug")),
			fmt.Sprintf(format, args...))...)
	}
}

func (l *std) Info(format string, args ...interface{}) {
	if InfoLevel <= l.lvl {
		log.Println(append(KeyValues(append(l.kvs, l.lvlKey, "info")),
			fmt.Sprintf(format, args...))...)
	}
}

func (l *std) Warn(err error) {
	if WarnLevel <= l.lvl {
		log.Println(append(KeyValues(append(l.kvs, l.lvlKey, "warn")),
			fmt.Sprintf("%v", err))...)
	}
}

func (l *std) Error(err error) {
	if ErrorLevel <= l.lvl {
		log.Println(append(KeyValues(append(l.kvs, l.lvlKey, "error")),
			fmt.Sprintf("%+v", err))...)
	}
}

func (l *std) Fatal(err error) {
	if FatalLevel <= l.lvl {
		log.Fatal(err)
	}
}

func (l *std) Panic(err error) {
	if PanicLevel < l.lvl {
		log.Panic(err)
	}
}

func KeyValues(kvs []interface{}) (vs []interface{}) {
	if len(kvs)%2 != 0 {
		return
	}

	for i := 0; i < len(kvs); i += 2 {
		vs = append(vs, fmt.Sprintf("%v=%v", kvs[i], kvs[i+1]))
	}

	return
}
