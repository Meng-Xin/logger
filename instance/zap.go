package logger

import (
	"context"
	"fmt"
	"github.com/Meng-Xin/logger"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Options func(config *ZapConfig)

type ZapConfig struct {
	// 日志配置
	Leave      string
	FilePath   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	// 上下文配置
	Service string
}

// defaultZapConfig 默认实例配置
func defaultZapConfig() ZapConfig {
	config := ZapConfig{
		MaxSize:    5,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
		Service:    "default_service",
	}
	config.FilePath = fmt.Sprintf("./logs/%s/%s.log", config.Service, config.Service)
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

func WithFilePath(filePath string) Options {
	return func(cfg *ZapConfig) {
		cfg.FilePath = filePath
	}
}

func WithMaxSize(maxSize int) Options {
	return func(cfg *ZapConfig) {
		cfg.MaxSize = maxSize
	}
}

func WithMaxBackups(maxBackups int) Options {
	return func(cfg *ZapConfig) {
		cfg.MaxBackups = maxBackups
	}
}

func WithMaxAge(age int) Options {
	return func(cfg *ZapConfig) {
		cfg.MaxAge = age
	}
}

func WithCompress(compress bool) Options {
	return func(cfg *ZapConfig) {
		cfg.Compress = compress
	}
}

func WithService(service string) Options {
	return func(cfg *ZapConfig) {
		cfg.Service = service
	}
}

// zapCenter 自定义实现类，实现ILog
type zapCenter struct {
	logger  *zap.Logger
	sugared *zap.SugaredLogger
}

func NewZapLogCenter(config *ZapConfig) logger.ILog {
	var coreArr []zapcore.Core
	//获取编码器
	encoderConfig := zap.NewProductionEncoderConfig()            //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder        //显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	//日志级别 [Debug,Error]
	allPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= getLogLeave(config.Leave)
	})

	//info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.FilePath,   //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    config.MaxSize,    //文件大小限制,单位MB
		MaxBackups: config.MaxBackups, //最大保留日志文件数量
		MaxAge:     config.MaxAge,     //日志文件保留天数
		Compress:   config.Compress,   //是否压缩处理
	})
	//第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	allFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), allPriority)

	coreArr = append(coreArr, allFileCore)
	//zap.AddCaller()为显示文件名和行号，可省略
	log := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())

	return &zapCenter{logger: log, sugared: log.Sugar()}
}

func (z *zapCenter) Debug(args ...any) {
	z.sugared.Debug(args...)
}

func (z *zapCenter) Info(args ...any) {
	z.sugared.Info(args...)
}

func (z *zapCenter) Warn(args ...any) {
	z.sugared.Warn(args...)
}

func (z *zapCenter) Error(args ...any) {
	z.sugared.Error(args...)
}

func (z *zapCenter) Fatal(args ...any) {
	z.sugared.Fatal(args...)
}

func (z *zapCenter) DebugContext(ctx context.Context, format string, args ...any) {
	// 获取上下文trace信息
	traceFields := getTraceInfo(ctx)
	// 使用 SugaredLogger 格式化消息
	message := getMessage(format, args...)

	// 使用 Logger 记录带有 traceFields 的结构化日志
	z.logger.With(traceFields...).Debug(message)
}

func (z *zapCenter) InfoContext(ctx context.Context, format string, args ...any) {
	// 获取上下文trace信息
	traceFields := getTraceInfo(ctx)
	// 使用 SugaredLogger 格式化消息
	message := getMessage(format, args...)

	// 使用 Logger 记录带有 traceFields 的结构化日志
	z.logger.With(traceFields...).Info(message)
}

func (z *zapCenter) WarnContext(ctx context.Context, format string, args ...any) {
	// 获取上下文trace信息
	traceFields := getTraceInfo(ctx)
	// 使用 SugaredLogger 格式化消息
	message := getMessage(format, args...)

	// 使用 Logger 记录带有 traceFields 的结构化日志
	z.logger.With(traceFields...).Warn(message)
}

func (z *zapCenter) ErrContext(ctx context.Context, format string, args ...any) {
	// 获取上下文trace信息
	traceFields := getTraceInfo(ctx)
	// 使用 SugaredLogger 格式化消息
	message := getMessage(format, args...)

	// 使用 Logger 记录带有 traceFields 的结构化日志
	z.logger.With(traceFields...).Error(message)
}

func (z *zapCenter) FatalContext(ctx context.Context, format string, args ...any) {
	// 获取上下文trace信息
	traceFields := getTraceInfo(ctx)
	// 使用 SugaredLogger 格式化消息
	message := getMessage(format, args...)

	// 使用 Logger 记录带有 traceFields 的结构化日志
	z.logger.With(traceFields...).Fatal(message)
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
	if data := ctx.Value("TraceInfoKey"); data != nil {
		if traceInfo, ok := data.(logger.TraceInfo); ok {
			traceFields = []zap.Field{
				zap.String("instance_id", traceInfo.InstanceID),
				zap.String("service_name", traceInfo.ServiceName),
				zap.String("service_version", traceInfo.ServiceVersion),
				zap.String("service_host", traceInfo.ServiceHost),
				zap.String("call_type", traceInfo.CallType),
				zap.String("trace_id", traceInfo.TraceID),
				zap.String("request_path", traceInfo.RequestPath),
			}
		}
	}
	return traceFields
}

// getLogLeave 获取日志等级
func getLogLeave(leave logger.LogLeave) zapcore.Level {
	if leave == "" {
		return zapcore.DebugLevel
	}
	switch leave {
	case logger.Debug:
		return zapcore.DebugLevel
	case logger.Info:
		return zapcore.InfoLevel
	case logger.Warn:
		return zapcore.WarnLevel
	case logger.Error:
		return zapcore.ErrorLevel
	case logger.Fatal:
		return zapcore.FatalLevel
	}
	return zapcore.DebugLevel
}
