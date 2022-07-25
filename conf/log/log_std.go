package log

import (
	"context"
	"fmt"
	"log"
)

func Std() Logger { return &std{lvl: DebugLevel, lvlKey: "@lvl"} }

type std struct {
	lvl    Level
	lvlKey string
	spans  []string
	kvs    []interface{}
}

func (d *std) SetLevel(lvl Level) { d.lvl = lvl }

func (d *std) SetLevelKey(k string) { d.lvlKey = k }

func (d *std) WithValues(kvs ...interface{}) Logger {
	return &std{
		lvl:   d.lvl,
		spans: d.spans,
		kvs:   append(d.kvs, kvs...),
	}
}

func (d *std) Start(ctx context.Context, name string, kvs ...interface{}) (context.Context, Logger) {
	return ctx, &std{
		lvl:   d.lvl,
		spans: append(d.spans, name),
		kvs:   append(d.kvs, kvs...)}
}

func (d *std) End() {
	if len(d.spans) != 0 {
		d.spans = d.spans[0 : len(d.spans)-1]
	}
}

func (d *std) Trace(format string, args ...interface{}) {
	if TraceLevel <= d.lvl {
		log.Println(append(KeyValues(append(d.kvs, d.lvlKey, "trace")),
			fmt.Sprintf(format, args...))...)
	}
}

func (d *std) Debug(format string, args ...interface{}) {
	if DebugLevel <= d.lvl {
		log.Println(append(KeyValues(append(d.kvs, d.lvlKey, "debug")),
			fmt.Sprintf(format, args...))...)
	}
}

func (d *std) Info(format string, args ...interface{}) {
	if InfoLevel <= d.lvl {
		log.Println(append(KeyValues(append(d.kvs, d.lvlKey, "info")),
			fmt.Sprintf(format, args...))...)
	}
}

func (d *std) Warn(err error) {
	if WarnLevel <= d.lvl {
		log.Println(append(KeyValues(append(d.kvs, d.lvlKey, "warn")),
			fmt.Sprintf("%v", err))...)
	}
}

func (d *std) Error(err error) {
	if ErrorLevel <= d.lvl {
		log.Println(append(KeyValues(append(d.kvs, d.lvlKey, "error")),
			fmt.Sprintf("%+v", err))...)
	}
}

func (d *std) Fatal(err error) {
	if FatalLevel <= d.lvl {
		log.Fatal(err)
	}
}

func (d *std) Panic(err error) {
	if PanicLevel < d.lvl {
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
