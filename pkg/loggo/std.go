package loggo

import "context"

var std *Loggo = &Loggo{}

func StandardLogger() Logger {
	return std
}

func SetServiceName(name string) {
	std.ServiceName = name
}

func SetBackend(backend Backend) {
	std.Backend = backend
}

func Fatal(ctx context.Context, params ...interface{}) {
	std.Fatal(ctx, params...)
}

func Error(ctx context.Context, params ...interface{}) {
	std.Error(ctx, params...)
}

func Warn(ctx context.Context, params ...interface{}) {
	std.Warn(ctx, params...)
}

func Info(ctx context.Context, params ...interface{}) {
	std.Info(ctx, params...)
}

func Debug(ctx context.Context, params ...interface{}) {
	std.Debug(ctx, params...)
}

func Trace(ctx context.Context, params ...interface{}) {
	std.Trace(ctx, params...)
}

func Fallback(ctx context.Context, params ...interface{}) {
	std.Fallback(ctx, params...)
}

func CircuitBreakerOpen(ctx context.Context, params ...interface{}) {
	std.CircuitBreakerOpen(ctx, params...)
}
