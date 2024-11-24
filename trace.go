package logger

type ContextLogCenterKey = string

const (
	LogTraceInfoKey ContextLogCenterKey = "context_log_trace_key"
)

// TraceInfo 日志链路追踪信息
type TraceInfo struct {
	InstanceID     string `json:"instance_id"`     //实例唯一id
	ServiceName    string `json:"service_name"`    //服务名
	ServiceVersion string `json:"service_version"` //服务版本号
	ServiceHost    string `json:"service_host"`    //服务主机地址
	CallType       string `json:"call_type"`       //调用类型:HTTP;GRPC;TRPC;TCP;UDP
	TraceID        string `json:"trace_id"`        //追踪id:请求唯一id
	RequestPath    string `json:"request_path"`    //请求路径
}
