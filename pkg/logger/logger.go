package logger

import (
	"os"

	"github.com/longln/go-ecommerce-backend-api/pkg/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)



type LoggerZap struct {
	*zap.Logger
}


func NewLogger(config setting.LogSetting) *LoggerZap{
	logLevel := "DEBUG"
	var level zapcore.Level
	switch logLevel {
	case "DEBUG":
		level = zap.DebugLevel
	case "INFO":
		level = zap.InfoLevel
	case "WARN":
		level = zap.WarnLevel
	case "ERROR":
		level = zap.ErrorLevel
	case "DPANIC":
		level = zap.DPanicLevel
	case "PANIC":
		level = zap.PanicLevel
	case "FATAL":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.FileLogName,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	core := zapcore.NewCore(encoder, 
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}




// format log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encodeConfig.EncodeCaller = zapcore.FullCallerEncoder

	return zapcore.NewConsoleEncoder(encodeConfig)
}
