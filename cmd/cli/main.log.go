package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))

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

// write log
func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE | os.O_RDWR, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	// ad console to print
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}