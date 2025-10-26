package logger

import "context"

// ILog 日志抽象层
type ILog interface {
	// Debug 无链路关系-debug
	Debug(args ...any)
	// Info 无链路关系-Info
	Info(args ...any)
	// Warn 无链路关系-Warn
	Warn(args ...any)
	// Error 无链路关系-Error
	Error(args ...any)
	// Fatal 无链路关系-Fatal
	Fatal(args ...any)

	// Debugf 无链路关系-debug
	Debugf(format string, args ...any)
	// Infof 无链路关系-Info
	Infof(format string, args ...any)
	// Warnf 无链路关系-Warn
	Warnf(format string, args ...any)
	// Errorf 无链路关系-Error
	Errorf(format string, args ...any)
	// Fatalf 无链路关系-Fatal
	Fatalf(format string, args ...any)

	// DebugContext 使用ctx传递上下文，使用日志链路追踪需要使用该方法
	DebugContext(ctx context.Context, format string, args ...any)
	// InfoContext 使用ctx传递上下文，使用日志链路追踪需要使用该方法
	InfoContext(ctx context.Context, format string, args ...any)
	// WarnContext 使用ctx传递上下文，使用日志链路追踪需要使用该方法
	WarnContext(ctx context.Context, format string, args ...any)
	// ErrContext 使用ctx传递上下文，使用日志链路追踪需要使用该方法
	ErrContext(ctx context.Context, format string, args ...any)
	// FatalContext 使用ctx传递上下文，使用日志链路追踪需要使用该方法
	FatalContext(ctx context.Context, format string, args ...any)
}
