package global

import (
	"context"

	"github.com/saitofun/qkit/conf/log"
	"github.com/saitofun/qkit/kit/sqlx"
	"github.com/saitofun/qkit/x/contextx"
)

type keyDatabase struct{}

func WithDatabase(db sqlx.DBExecutor) contextx.WithContext {
	return func(ctx context.Context) context.Context {
		return contextx.WithValue(ctx, keyDatabase{}, db)
	}
}

func DBExecutorFromContext(ctx context.Context) sqlx.DBExecutor {
	return ctx.Value(keyDatabase{}).(sqlx.DBExecutor).WithContext(ctx)
}

type keyLogger struct{}

func WithLogger(l log.Logger) contextx.WithContext {
	return func(ctx context.Context) context.Context {
		return contextx.WithValue(ctx, keyLogger{}, l)
	}
}

func LoggerFromContext(ctx context.Context) log.Logger {
	return ctx.Value(keyLogger{}).(log.Logger)
}
