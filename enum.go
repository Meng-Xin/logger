package logger

type LogLeave = string

const (
	Debug LogLeave = "debug"
	Info  LogLeave = "info"
	Warn  LogLeave = "warn"
	Error LogLeave = "error"
	Fatal LogLeave = "fatal"
)
