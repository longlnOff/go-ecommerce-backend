package logger

import (
	"os"

	"github.com/longln/go-ecommerce-backend/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)


type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.LogLevel
	// debug --> info --> warning --> error --> fatal --> panic
	var level zapcore.Level
	switch logLevel {
		case "debug":
			level = zap.DebugLevel
		case "info":
			level = zap.InfoLevel
		case "warning":
			level = zap.WarnLevel
		case "error":
			level = zap.ErrorLevel
		case "fatal":
			level = zap.FatalLevel
		case "panic":
			level = zap.PanicLevel
		default:
			level = zap.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename: config.FileLogName,
		MaxSize: config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge: config.MaxAge,
		Compress: config.Compress,
	}
	core := zapcore.NewCore(encoder, 
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), 
		level)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	return &LoggerZap{
		Logger: logger,
	}
}

func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// time format  
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts --> time
	encoderConfig.TimeKey = "time"

	// uppercase INFOR, nifor --> INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// print caller (trace)
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}