package log_test

import (
	"context"

	. "github.com/saitofun/qkit/conf/log"
)

func ExampleLogger() {
	ctx := WithLogger(context.Background(), Std())

	l := FromContext(ctx).WithValues("k", "k")

	l.Debug("test %d", 1)
	l.Trace("test %d", 1)
	l.Info("test %d", 1)
	// Output:
}

func ExampleLogger_Start() {
	ctx := WithLogger(context.Background(), Std())

	_, l := Start(ctx, "span", "k", "k")
	defer l.End()

	l.Debug("test %d", 1)
	l.Trace("test %d", 1)
	l.Info("test %d", 1)
	// Output:
}
