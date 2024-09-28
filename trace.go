package logger

type ContextKey = string

const (
	TraceInfoKey ContextKey = "context_log_trace_key"
)

// TraceInfo 日志链路追踪信息
type TraceInfo struct {
	// 服务实例相关
	InstanceID     string
	ServiceName    string
	ServiceVersion string
	ServiceHost    string
	// 服务请求配置
	CallType    string
	TraceID     string
	RequestPath string
}
