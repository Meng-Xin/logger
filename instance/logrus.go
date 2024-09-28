package logger

import (
	"context"
	"fmt"
	"github.com/Meng-Xin/logger"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type logrusCenter struct {
	logger *logrus.Logger
}

// NewLogrusCenter 新建logrus日志中心
func NewLogrusCenter(level string, filePath string, hooks ...logrus.Hook) logger.ILog {
	parseLevel, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err.Error())
	}
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile " + filePath)
		panic(err)
	}
	log := &logrus.Logger{
		Out:   io.MultiWriter(f, os.Stdout),         // 文件 + 控制台输出
		Level: parseLevel,                           // 日志等级
		Hooks: make(map[logrus.Level][]logrus.Hook), // 初始化Hook Map,否则导致Hook添加过程中的空指针引用。
		Formatter: &logrus.JSONFormatter{ // Json形式输出
			TimestampFormat: "2006-01-02 15:04:05", //日期格式
		},
	}
	// 绑定Hook
	for index, _ := range hooks {
		log.AddHook(hooks[index])
	}
	log.Infof("日志服务启动成功")
	return &logrusCenter{logger: log}
}

func (l *logrusCenter) Debug(args ...any) {
	l.logger.Debug(args...)
}

func (l *logrusCenter) Info(args ...any) {
	l.logger.Info(args...)
}

func (l *logrusCenter) Warn(args ...any) {
	l.logger.Warn(args...)
}

func (l *logrusCenter) Error(args ...any) {
	l.logger.Error(args...)
}

func (l *logrusCenter) Fatal(args ...any) {
	l.logger.Fatal(args...)
}

func (l *logrusCenter) DebugContext(ctx context.Context, format string, args ...any) {
	l.logger.WithContext(ctx).Debugf(format, args...)
}

func (l *logrusCenter) InfoContext(ctx context.Context, format string, args ...any) {
	l.logger.WithContext(ctx).Infof(format, args...)
}

func (l *logrusCenter) WarnContext(ctx context.Context, format string, args ...any) {
	l.logger.WithContext(ctx).Warnf(format, args...)
}

func (l *logrusCenter) ErrContext(ctx context.Context, format string, args ...any) {
	l.logger.WithContext(ctx).Errorf(format, args...)
}

func (l *logrusCenter) FatalContext(ctx context.Context, format string, args ...any) {
	l.logger.WithContext(ctx).Fatalf(format, args...)
}
