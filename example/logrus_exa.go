package example

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func NewLogrusCenter(level string, filePath string, hooks ...logrus.Hook) *logrus.Logger {
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
	return log
}
