# Go Trace Logger

一个基于[zap](https://github.com/uber-go/zap)的高性能日志库，支持链路追踪功能。

## 特性

- 基于高性能的zap日志库
- 支持链路追踪，轻松实现分布式系统的日志追踪
- 支持日志文件轮转
- 支持多种日志级别
- 支持JSON格式输出
- 自动记录调用位置
- 支持丰富的配置选项

## 安装

```bash
go get github.com/Meng-Xin/logger
```

## 快速开始

### 基本使用

```go
package main

import (
    "github.com/Meng-Xin/logger"
)

func main() {
    // 创建日志配置
    config := logger.NewZapConfig(
        logger.WithServiceName("my-service"),
    )
    
    // 创建日志实例
    log := logger.NewZapLogCenter(config)
    
    // 记录日志
    log.Info("Hello, Logger!")
    log.Error("Something went wrong")
}
```

### 使用链路追踪

```go
package main

import (
    "context"
    "github.com/Meng-Xin/logger"
)

func main() {
    // 创建日志配置
    config := logger.NewZapConfig(
        logger.WithServiceName("my-service"),
    )
    
    // 创建日志实例
    log := logger.NewZapLogCenter(config)
    
    // 创建带有追踪信息的上下文
    traceInfo := logger.TraceInfo{
        InstanceID:     "instance-001",
        ServiceName:    "my-service",
        ServiceVersion: "v1.0.0",
        ServiceHost:    "localhost",
        CallType:       "HTTP",
        TraceID:        "trace-123",
        SpanID:         "span-456",
        RequestPath:    "/api/v1/users",
    }
    
    ctx := context.WithValue(context.Background(), logger.LogTraceInfoKey, traceInfo)
    
    // 记录带追踪信息的日志
    log.InfoContext(ctx, "Processing request for user %s", "john")
}
```

## 配置选项

### 日志配置

```go
config := logger.NewZapConfig(
    // 设置服务名称（影响日志文件路径）
    logger.WithServiceName("my-service"),
    
    // 设置单个日志文件的最大大小（MB）
    logger.WithMaxSize(10),
    
    // 设置最大日志文件数量
    logger.WithMaxBackups(5),
    
    // 设置日志文件保留天数
    logger.WithMaxAge(30),
    
    // 设置是否压缩旧日志文件
    logger.WithCompress(true),
)
```

### 日志级别

支持以下日志级别：
- Debug
- Info
- Warn
- Error
- Fatal

### 链路追踪信息

TraceInfo结构包含以下字段：
- InstanceID：实例唯一标识
- ServiceName：服务名称
- ServiceVersion：服务版本号
- ServiceHost：服务主机地址
- CallType：调用类型（HTTP/GRPC/TRPC/TCP/UDP）
- TraceID：请求唯一标识
- SpanID：追踪单元标识
- RequestPath：请求路径

## 日志输出示例

```json
{
  "level": "INFO",
  "timestamp": "2023-05-20T10:30:00.000Z",
  "caller": "app/handler.go:42",
  "msg": "Processing request for user john",
  "instance_id": "instance-001",
  "service_name": "my-service",
  "service_version": "v1.0.0",
  "service_host": "localhost",
  "call_type": "HTTP",
  "trace_id": "trace-123",
  "span_id": "span-456",
  "request_path": "/api/v1/users"
}
```

## 最佳实践

1. 总是使用有意义的服务名称，这将帮助你更好地组织日志文件
2. 在分布式系统中，确保为每个请求生成唯一的TraceID
3. 合理配置日志轮转参数，避免磁盘空间耗尽
4. 在关键节点使用链路追踪功能，帮助排查问题
5. 根据环境选择适当的日志级别

## 性能考虑

- 该日志库基于高性能的zap日志库
- JSON编码器提供了结构化的日志输出
- 支持日志文件轮转，避免单个文件过大
- 异步写入，不会阻塞主程序执行

## 贡献

欢迎提交Issue和Pull Request！

## 许可证

MIT License
