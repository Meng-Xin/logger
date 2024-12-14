package logger

import (
	"github.com/Meng-Xin/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

func GinMiddleware(log logger.ILog) gin.HandlerFunc {
	logTrace := logger.TraceInfo{
		InstanceID:     "",
		ServiceName:    "default",
		ServiceVersion: "1.0.0",
		ServiceHost:    logger.GetLocalIP(),
		CallType:       "HTTP",
		TraceID:        "",
		SpanID:         "",
		RequestPath:    "",
	}

	return func(c *gin.Context) {
		// 请求路径
		logTrace.RequestPath = c.FullPath()

		// 链路追踪信息
		span := trace.SpanFromContext(c.Request.Context())
		spanCtx := span.SpanContext()

		if spanCtx.IsValid() {
			logTrace.TraceID = spanCtx.TraceID().String()
			logTrace.SpanID = spanCtx.SpanID().String()
		} else if c.Param("trace_id") != "" {
			logTrace.TraceID = c.Param("trace_id")
		} else {
			logTrace.TraceID = uuid.New().String()
		}
		c.Set(logger.LogTraceInfoKey, logTrace)
		// 入口日志
		log.DebugContext(c, "======"+logTrace.RequestPath+"======"+"start")
		c.Next()
		// 出口日志
		log.DebugContext(c, "======"+logTrace.RequestPath+"======"+"end")
	}
}
