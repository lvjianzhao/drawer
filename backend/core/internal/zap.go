package internal

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
	"time"

	"server/global"
)

type _zap struct{}

var Zap = new(_zap)

// CustomTimeEncoder 自定义日志输出时间格式
func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.TD27_CONFIG.Zap.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

func CustomCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	// 获取完整的调用者路径, eg:  D:/code/vue/drawer/backend/service/customer/customer.go:16
	fullPath := caller.FullPath()
	// 以项目工程目录的名称作为截断, eg: [D:/code/vue/drawer/ /service/customer/customer.go:16]
	parts := strings.Split(fullPath, "backend")
	lastPart := parts[len(parts)-1]
	lastPartStr := fmt.Sprintf("%s", lastPart)
	enc.AppendString(lastPartStr)
}

// GetEncoderConfig 获取zapcore.EncoderConfig
func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.TD27_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    global.TD27_CONFIG.Zap.ZapEncodeLevel(),
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   CustomCallerEncoder, // zapcore.FullCallerEncoder显示绝对路径,zapcore.ShortCallerEncoder显示相对路径
	}
}

// GetEncoder 获取 zapcore.Encoder
func (z *_zap) GetEncoder() zapcore.Encoder {
	if global.TD27_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

// GetEncoderCore 获取Encoder的 zapcore.Core
func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	// 使用lumberjack进行日志分割
	writer := LumberjackLogs.GetWriteSyncer(l.String())

	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}

// GetZapCores 根据配置文件的Level获取 []zapcore.Core
func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := global.TD27_CONFIG.Zap.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}
