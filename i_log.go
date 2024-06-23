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

	// DebugWithContext 存在链路关系-debug
	DebugWithContext(ctx context.Context, args ...any)
	// InfoWithContext 存在链路关系-info
	InfoWithContext(ctx context.Context, args ...any)
	// WarnWithContext 存在链路关系-war
	WarnWithContext(ctx context.Context, args ...any)
	// ErrorWithContext 存在链路关系-error
	ErrorWithContext(ctx context.Context, args ...any)
	// FatalWithContext 存在链路关系-fatal
	FatalWithContext(ctx context.Context, args ...any)
}
