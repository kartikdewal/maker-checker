package logger

import (
	"context"
)

type ContextLogger interface {
	With(args ...interface{}) ContextLogger
	WithoutContext() Logger
	Named(n string) ContextLogger
	Debug(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Warn(ctx context.Context, args ...interface{})

	Error(ctx context.Context, args ...interface{})

	Panic(ctx context.Context, args ...interface{})

	Fatal(ctx context.Context, args ...interface{})

	Debugf(ctx context.Context, template string, args ...interface{})

	Infof(ctx context.Context, template string, args ...interface{})

	Warnf(ctx context.Context, template string, args ...interface{})

	Errorf(ctx context.Context, template string, args ...interface{})

	Panicf(ctx context.Context, template string, args ...interface{})

	Fatalf(ctx context.Context, template string, args ...interface{})

	Debugw(ctx context.Context, msg string, keysAndValues ...interface{})

	Infow(ctx context.Context, msg string, keysAndValues ...interface{})
	Warnw(ctx context.Context, msg string, keysAndValues ...interface{})

	Errorw(ctx context.Context, msg string, keysAndValues ...interface{})

	Panicw(ctx context.Context, msg string, keysAndValues ...interface{})

	Fatalw(ctx context.Context, msg string, keysAndValues ...interface{})
	Sync() error
}
