package logger

import "context"

type NoOpLogger struct {
}

func (n NoOpLogger) Named(_ string) Logger {
	return n
}

func (n NoOpLogger) With(args ...interface{}) Logger {
	return n
}

func (n NoOpLogger) Debug(args ...interface{}) {
}

func (n NoOpLogger) Info(args ...interface{}) {
}

func (n NoOpLogger) Warn(args ...interface{}) {
}

func (n NoOpLogger) Error(args ...interface{}) {
}

func (n NoOpLogger) Panic(args ...interface{}) {
}

func (n NoOpLogger) Fatal(args ...interface{}) {
}

func (n NoOpLogger) Debugf(template string, args ...interface{}) {
}

func (n NoOpLogger) Infof(template string, args ...interface{}) {
}

func (n NoOpLogger) Warnf(template string, args ...interface{}) {
}

func (n NoOpLogger) Errorf(template string, args ...interface{}) {
}

func (n NoOpLogger) Panicf(template string, args ...interface{}) {
}

func (n NoOpLogger) Fatalf(template string, args ...interface{}) {
}

func (n NoOpLogger) Debugw(msg string, keysAndValues ...interface{}) {
}

func (n NoOpLogger) Infow(msg string, keysAndValues ...interface{}) {
}

func (n NoOpLogger) Warnw(msg string, keysAndValues ...interface{}) {
}

func (n NoOpLogger) Errorw(msg string, keysAndValues ...interface{}) {
}

func (n NoOpLogger) Panicw(msg string, keysAndValues ...interface{}) {
}

func (n NoOpLogger) Fatalw(msg string, keysAndValues ...interface{}) {
}

func (n NoOpLogger) Sync() error {
	return nil
}

type NoOpContextLogger struct {
}

func (n NoOpContextLogger) WithoutContext() Logger {
	return NoOpLogger{}
}

func (n NoOpContextLogger) With(args ...interface{}) ContextLogger {
	return n
}

func (n NoOpContextLogger) Named(_ string) ContextLogger {
	return n
}

func (n NoOpContextLogger) Debug(ctx context.Context, args ...interface{}) {
}

func (n NoOpContextLogger) Info(ctx context.Context, args ...interface{}) {
}

func (n NoOpContextLogger) Warn(ctx context.Context, args ...interface{}) {

}

func (n NoOpContextLogger) Error(ctx context.Context, args ...interface{}) {

}

func (n NoOpContextLogger) Panic(ctx context.Context, args ...interface{}) {

}

func (n NoOpContextLogger) Fatal(ctx context.Context, args ...interface{}) {

}

func (n NoOpContextLogger) Debugf(ctx context.Context, template string, args ...interface{}) {

}

func (n NoOpContextLogger) Infof(ctx context.Context, template string, args ...interface{}) {

}

func (n NoOpContextLogger) Warnf(ctx context.Context, template string, args ...interface{}) {

}

func (n NoOpContextLogger) Errorf(ctx context.Context, template string, args ...interface{}) {

}

func (n NoOpContextLogger) Panicf(ctx context.Context, template string, args ...interface{}) {

}

func (n NoOpContextLogger) Fatalf(ctx context.Context, template string, args ...interface{}) {

}

func (n NoOpContextLogger) Debugw(ctx context.Context, msg string, keysAndValues ...interface{}) {

}

func (n NoOpContextLogger) Infow(ctx context.Context, msg string, keysAndValues ...interface{}) {

}

func (n NoOpContextLogger) Warnw(ctx context.Context, msg string, keysAndValues ...interface{}) {

}

func (n NoOpContextLogger) Errorw(ctx context.Context, msg string, keysAndValues ...interface{}) {

}

func (n NoOpContextLogger) Panicw(ctx context.Context, msg string, keysAndValues ...interface{}) {

}

func (n NoOpContextLogger) Fatalw(ctx context.Context, msg string, keysAndValues ...interface{}) {

}

func (n NoOpContextLogger) Sync() error {
	return nil
}
