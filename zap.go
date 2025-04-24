package logger

import (
	"context"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Options func(config *ZapConfig)

// ZapConfig zap 日志配置类
type ZapConfig struct {
	Leave       string //日志等级
	ServiceName string //服务名
	MaxSize     int    //单文件最大单位MB
	MaxBackups  int    //最大分割数
	MaxAge      int    //最大保存时间
	Compress    bool   //是否压缩处理
	FilePath    string //日志路径
}

// defaultZapConfig 默认实例配置
func defaultZapConfig() ZapConfig {
	config := ZapConfig{
		Leave:      Debug,
		FilePath:   DefaultFilePath,
		MaxSize:    5,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return config
}

// NewZapConfig 创建Zap日志配置
func NewZapConfig(options ...Options) *ZapConfig {
	cfg := defaultZapConfig()
	for _, option := range options {
		option(&cfg)
	}
	return &cfg
}

// WithServiceName 设置文件路径-需传入对应服务名
// 例如：hello -> 日志路径 /logger/logs/hello/hello.log
func WithServiceName(serviceName string) Options {
	return func(cfg *ZapConfig) {
		cfg.ServiceName = serviceName
		cfg.FilePath = fmt.Sprintf("./logs/%s/%s.log", serviceName, serviceName)
	}
}

// WithMaxSize 设置文件最大单位 MB
func WithMaxSize(maxSize int) Options {
	return func(cfg *ZapConfig) {
		cfg.MaxSize = maxSize
	}
}

// WithMaxBackups 设置最大分割数
func WithMaxBackups(maxBackups int) Options {
	return func(cfg *ZapConfig) {
		cfg.MaxBackups = maxBackups
	}
}

// WithMaxAge 设置最大保存时间
func WithMaxAge(age int) Options {
	return func(cfg *ZapConfig) {
		cfg.MaxAge = age
	}
}

// WithCompress 设置是否压缩处理
func WithCompress(compress bool) Options {
	return func(cfg *ZapConfig) {
		cfg.Compress = compress
	}
}

// zapCenter 自定义实现类，实现ILog
type zapCenter struct {
	logger  *zap.Logger
	sugared *zap.SugaredLogger
}

func NewZapLogCenter(config *ZapConfig) ILog {
	var coreArr []zapcore.Core
	//获取编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder //FullCallerEncoder显示完整文件路径
	encoder := zapcore.NewJSONEncoder(encoderConfig)        //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式

	//日志级别 [Debug,Error]
	allPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= getLogLeave(config.Leave)
	})

	//日志本地写入
	localFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.FilePath,   //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    config.MaxSize,    //文件大小限制,单位MB
		MaxBackups: config.MaxBackups, //最大保留日志文件数量
		MaxAge:     config.MaxAge,     //日志文件保留天数
		Compress:   config.Compress,   //是否压缩处理
	})
	//第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	allFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(localFileWriteSyncer, zapcore.AddSync(os.Stdout)), allPriority)
	coreArr = append(coreArr, allFileCore)
	//zap.AddCaller()为显示文件名和行号，可省略
	log := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())

	return &zapCenter{logger: log, sugared: log.Sugar()}
}

func (z *zapCenter) Debug(args ...any) {
	z.sugared.WithOptions(zap.AddCallerSkip(1)).Debug(args...)
}

func (z *zapCenter) Info(args ...any) {
	z.sugared.WithOptions(zap.AddCallerSkip(1)).Info(args...)
}

func (z *zapCenter) Warn(args ...any) {
	z.sugared.WithOptions(zap.AddCallerSkip(1)).Warn(args...)
}

func (z *zapCenter) Error(args ...any) {
	z.sugared.WithOptions(zap.AddCallerSkip(1)).Error(args...)
}

func (z *zapCenter) Fatal(args ...any) {
	z.sugared.WithOptions(zap.AddCallerSkip(1)).Fatal(args...)
}

func (z *zapCenter) DebugContext(ctx context.Context, format string, args ...any) {
	// 获取上下文trace信息
	traceFields := getTraceInfo(ctx)
	// 使用 SugaredLogger 格式化消息
	message := getMessage(format, args...)

	// 使用 Logger 记录带有 traceFields 的结构化日志
	z.logger.WithOptions(zap.AddCallerSkip(1)).With(traceFields...).Debug(message)
}

func (z *zapCenter) InfoContext(ctx context.Context, format string, args ...any) {
	// 获取上下文trace信息
	traceFields := getTraceInfo(ctx)
	// 使用 SugaredLogger 格式化消息
	message := getMessage(format, args...)

	// 使用 Logger 记录带有 traceFields 的结构化日志
	z.logger.WithOptions(zap.AddCallerSkip(1)).With(traceFields...).Info(message)
}

func (z *zapCenter) WarnContext(ctx context.Context, format string, args ...any) {
	// 获取上下文trace信息
	traceFields := getTraceInfo(ctx)
	// 使用 SugaredLogger 格式化消息
	message := getMessage(format, args...)

	// 使用 Logger 记录带有 traceFields 的结构化日志
	z.logger.WithOptions(zap.AddCallerSkip(1)).With(traceFields...).Warn(message)
}

func (z *zapCenter) ErrContext(ctx context.Context, format string, args ...any) {
	// 获取上下文trace信息
	traceFields := getTraceInfo(ctx)
	// 使用 SugaredLogger 格式化消息
	message := getMessage(format, args...)

	// 使用 Logger 记录带有 traceFields 的结构化日志
	z.logger.WithOptions(zap.AddCallerSkip(1)).With(traceFields...).Error(message)
}

func (z *zapCenter) FatalContext(ctx context.Context, format string, args ...any) {
	// 获取上下文trace信息
	traceFields := getTraceInfo(ctx)
	// 使用 SugaredLogger 格式化消息
	message := getMessage(format, args...)

	// 使用 Logger 记录带有 traceFields 的结构化日志
	z.logger.WithOptions(zap.AddCallerSkip(1)).With(traceFields...).Fatal(message)
}

// getMessage format with Sprint, Sprintf, or neither.
func getMessage(template string, fmtArgs ...any) string {
	if len(fmtArgs) == 0 {
		return template
	}

	if template != "" {
		return fmt.Sprintf(template, fmtArgs...)
	}

	if len(fmtArgs) == 1 {
		if str, ok := fmtArgs[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(fmtArgs...)
}

// getTraceInfo 拿到日志上下文
func getTraceInfo(ctx context.Context) (traceFields []zap.Field) {
	if data := ctx.Value(LogTraceInfoKey); data != nil {
		if traceInfo, ok := data.(TraceInfo); ok {
			traceFields = []zap.Field{
				zap.String("request_path", traceInfo.RequestPath),
				zap.String("trace_id", traceInfo.TraceID),
				zap.String("service_name", traceInfo.ServiceName),
				zap.String("call_type", traceInfo.CallType),
				zap.String("instance_id", traceInfo.InstanceID),
				zap.String("service_version", traceInfo.ServiceVersion),
				zap.String("service_host", traceInfo.ServiceHost),
			}
		}
	}
	return traceFields
}

// getLogLeave 获取日志等级
func getLogLeave(leave LogLeave) zapcore.Level {
	if leave == "" {
		return zapcore.DebugLevel
	}
	switch leave {
	case Debug:
		return zapcore.DebugLevel
	case Info:
		return zapcore.InfoLevel
	case Warn:
		return zapcore.WarnLevel
	case Error:
		return zapcore.ErrorLevel
	case Fatal:
		return zapcore.FatalLevel
	}
	return zapcore.DebugLevel
}
