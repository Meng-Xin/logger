package logger

import (
	"context"
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestNewZapLogCenter(t *testing.T) {
	config := NewZapConfig(
		WithServiceName("test"),
		WithMaxSize(1),
		WithMaxBackups(1),
		WithMaxAge(1),
		WithCompress(false),
	)
	logCenter := NewZapLogCenter(config)

	// 测试 Debug 方法
	logCenter.Debug("debug message")

	// 测试 Info 方法
	logCenter.Info("info message")

	// 测试 Warn 方法
	logCenter.Warn("warn message")

	// 测试 Error 方法
	logCenter.Error("error message")

	// 测试 Fatal 方法
	// logCenter.Fatal("fatal message") // 注意：运行此测试时会导致程序退出，可以注释掉这行代码

	// 测试 DebugContext 方法
	ctx := context.WithValue(context.Background(), LogTraceInfoKey, TraceInfo{
		InstanceID:     "",
		ServiceName:    "test",
		ServiceVersion: "1.0.0",
		ServiceHost:    "127.0.0.1:8080",
		CallType:       "grpc",
		TraceID:        "b98f757b-e9b9-4e0c-a8b2-609c5cbcf990",
		RequestPath:    "/test",
	})
	logCenter.DebugContext(ctx, "debug context message")

	// 测试 InfoContext 方法
	logCenter.InfoContext(ctx, "info context message")

	// 测试 WarnContext 方法
	logCenter.WarnContext(ctx, "warn context message")

	// 测试 ErrContext 方法
	logCenter.ErrContext(ctx, "error context message")

	// 测试 FatalContext 方法
	// logCenter.FatalContext(ctx, "fatal context message") // 注意：运行此测试时会导致程序退出，可以注释掉这行代码
}

func TestGetMessage(t *testing.T) {
	// 测试没有格式化参数的情况
	message := getMessage("plain message")
	if message != "plain message" {
		t.Errorf("Expected 'plain message', got '%s'", message)
	}

	// 测试有格式化参数的情况
	message = getMessage("formatted message: %d", 123)
	if message != "formatted message: 123" {
		t.Errorf("Expected 'formatted message: 123', got '%s'", message)
	}

	// 测试没有格式化参数的情况
	message = getMessage("plain message")
	if message != "plain message" {
		t.Errorf("Expected 'plain message', got '%s'", message)
	}

}

func TestGetTraceInfo(t *testing.T) {
	ctx := context.WithValue(context.Background(), "TraceInfoKey", TraceInfo{
		InstanceID:     "test_instance_id",
		ServiceName:    "test_service_name",
		ServiceVersion: "test_service_version",
		ServiceHost:    "test_service_host",
		CallType:       "test_call_type",
		TraceID:        "test_trace_id",
		RequestPath:    "test_request_path",
	})
	traceFields := getTraceInfo(ctx)
	if len(traceFields) != 7 {
		t.Errorf("Expected 7 trace fields, got %d", len(traceFields))
	}
}

func TestGetLogLeave(t *testing.T) {
	level := getLogLeave(Debug)
	if level != zapcore.DebugLevel {
		t.Errorf("Expected DebugLevel, got %v", level)
	}

	level = getLogLeave(Info)
	if level != zapcore.InfoLevel {
		t.Errorf("Expected InfoLevel, got %v", level)
	}

	level = getLogLeave(Warn)
	if level != zapcore.WarnLevel {
		t.Errorf("Expected WarnLevel, got %v", level)
	}

	level = getLogLeave(Error)
	if level != zapcore.ErrorLevel {
		t.Errorf("Expected ErrorLevel, got %v", level)
	}

	level = getLogLeave(Fatal)
	if level != zapcore.FatalLevel {
		t.Errorf("Expected FatalLevel, got %v", level)
	}

	level = getLogLeave("")
	if level != zapcore.DebugLevel {
		t.Errorf("Expected DebugLevel for empty leave, got %v", level)
	}
}
