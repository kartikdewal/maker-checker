package logger

import (
	"context"
	"go.uber.org/zap"
)

func NewContextLogger(log *zap.SugaredLogger) ContextLogger {
	return &contextLogger{log: log.WithOptions(zap.AddCallerSkip(1))}
}

type contextLogger struct {
	log *zap.SugaredLogger
}

func (z contextLogger) addSpan(ctx context.Context) *zap.SugaredLogger {

	l := z.log.With()

	return l
}

func (z contextLogger) With(args ...interface{}) ContextLogger {
	return contextLogger{log: z.log.With(args...)}
}

func (z contextLogger) Named(n string) ContextLogger {
	return contextLogger{log: z.log.Named(n)}
}

func (z contextLogger) Debug(ctx context.Context, args ...interface{}) {
	z.addSpan(ctx).Debug(args...)
}

func (z contextLogger) Info(ctx context.Context, args ...interface{}) {
	z.addSpan(ctx).Info(args...)
}

func (z contextLogger) Warn(ctx context.Context, args ...interface{}) {
	z.addSpan(ctx).Warn(args...)
}

func (z contextLogger) Error(ctx context.Context, args ...interface{}) {
	z.addSpan(ctx).Error(args...)
}

func (z contextLogger) Panic(ctx context.Context, args ...interface{}) {
	z.addSpan(ctx).Panic(args...)
}

func (z contextLogger) Fatal(ctx context.Context, args ...interface{}) {
	z.addSpan(ctx).Fatal(args...)
}

func (z contextLogger) Debugf(ctx context.Context, template string, args ...interface{}) {
	z.addSpan(ctx).Debugf(template, args...)
}

func (z contextLogger) Infof(ctx context.Context, template string, args ...interface{}) {
	z.addSpan(ctx).Infof(template, args...)
}

func (z contextLogger) Warnf(ctx context.Context, template string, args ...interface{}) {
	z.addSpan(ctx).Warnf(template, args...)
}

func (z contextLogger) Errorf(ctx context.Context, template string, args ...interface{}) {
	z.addSpan(ctx).Errorf(template, args...)
}

func (z contextLogger) Panicf(ctx context.Context, template string, args ...interface{}) {
	z.addSpan(ctx).Panicf(template, args...)
}

func (z contextLogger) Fatalf(ctx context.Context, template string, args ...interface{}) {
	z.addSpan(ctx).Fatalf(template, args...)
}

func (z contextLogger) Debugw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	z.addSpan(ctx).Debugw(msg, keysAndValues...)
}

func (z contextLogger) Infow(ctx context.Context, msg string, keysAndValues ...interface{}) {
	z.addSpan(ctx).Infow(msg, keysAndValues...)
}

func (z contextLogger) Warnw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	z.addSpan(ctx).Warnw(msg, keysAndValues...)
}

func (z contextLogger) Errorw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	z.addSpan(ctx).Errorw(msg, keysAndValues...)
}

func (z contextLogger) Panicw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	z.addSpan(ctx).Panicw(msg, keysAndValues...)
}

func (z contextLogger) Fatalw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	z.addSpan(ctx).Fatalw(msg, keysAndValues...)
}

func (z contextLogger) Sync() error {
	return z.log.Sync()
}

func (z contextLogger) WithoutContext() Logger {
	return &logProxy{sLogger: z.log}
}

// logProxy acts a proxy between the logger.Logger interface and the *zap.SugaredLogger.
type logProxy struct {
	sLogger *zap.SugaredLogger
}

func (p logProxy) Debug(args ...interface{}) {
	p.sLogger.Debug(args...)
}

func (p logProxy) Info(args ...interface{}) {
	p.sLogger.Info(args...)
}

func (p logProxy) Warn(args ...interface{}) {
	p.sLogger.Warn(args...)
}

func (p logProxy) Error(args ...interface{}) {
	p.sLogger.Error(args...)
}

func (p logProxy) Panic(args ...interface{}) {
	p.sLogger.Panic(args...)
}

func (p logProxy) Fatal(args ...interface{}) {
	p.sLogger.Fatal(args...)
}

func (p logProxy) Debugf(template string, args ...interface{}) {
	p.sLogger.Debugf(template, args...)
}

func (p logProxy) Infof(template string, args ...interface{}) {
	p.sLogger.Infof(template, args...)
}

func (p logProxy) Warnf(template string, args ...interface{}) {
	p.sLogger.Infof(template, args...)
}

func (p logProxy) Errorf(template string, args ...interface{}) {
	p.sLogger.Errorf(template, args...)
}

func (p logProxy) Panicf(template string, args ...interface{}) {
	p.sLogger.Panicf(template, args...)
}

func (p logProxy) Fatalf(template string, args ...interface{}) {
	p.sLogger.Fatalf(template, args...)
}

func (p logProxy) Debugw(msg string, keysAndValues ...interface{}) {
	p.sLogger.Debugw(msg, keysAndValues...)
}

func (p logProxy) Infow(msg string, keysAndValues ...interface{}) {
	p.sLogger.Infow(msg, keysAndValues...)
}

func (p logProxy) Warnw(msg string, keysAndValues ...interface{}) {
	p.sLogger.Warnw(msg, keysAndValues...)
}

func (p logProxy) Errorw(msg string, keysAndValues ...interface{}) {
	p.sLogger.Errorw(msg, keysAndValues...)
}

func (p logProxy) Panicw(msg string, keysAndValues ...interface{}) {
	p.sLogger.Panicw(msg, keysAndValues...)
}

func (p logProxy) Fatalw(msg string, keysAndValues ...interface{}) {
	p.sLogger.Fatalw(msg, keysAndValues...)
}

func (p logProxy) Sync() error {
	return p.sLogger.Sync()
}

func (p logProxy) Named(n string) Logger {
	return logProxy{
		sLogger: p.sLogger.Named(n),
	}
}

func (p logProxy) With(args ...interface{}) Logger {
	return logProxy{
		sLogger: p.sLogger.With(args...),
	}
}
