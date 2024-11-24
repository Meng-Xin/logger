package logger

// LogLeave 日志等级
type LogLeave = string

const (
	Debug LogLeave = "debug"
	Info  LogLeave = "info"
	Warn  LogLeave = "warn"
	Error LogLeave = "error"
	Fatal LogLeave = "fatal"
)

// DefaultFilePath 日志中心默认日志路径
const DefaultFilePath = "../logs/default.log"
