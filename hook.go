package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// emailLoggerHook 邮件日志hook
type emailLoggerHook struct {
}

func NewEmail() logrus.Hook {
	return &emailLoggerHook{}
}

// Levels 需要监控的日志等级，只有命中列表中的日志等级才会触发Hook
func (l *emailLoggerHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
	}
}

// Fire 触发钩子函数，本实例为触发后发送邮件报警。
func (l *emailLoggerHook) Fire(entry *logrus.Entry) error {
	// 触发loggerHook函数调用
	fmt.Println("触发loggerHook函数调用")
	return nil
}

// kafkaLoggerHook kafka消息队列日志hook
type kafkaLoggerHook struct{}

func NewKafkaLoggerHook() logrus.Hook {
	return &kafkaLoggerHook{}
}
func (k *kafkaLoggerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (k *kafkaLoggerHook) Fire(entry *logrus.Entry) error {
	// 调用消息队列获取消息
	panic("not implemented")
}
