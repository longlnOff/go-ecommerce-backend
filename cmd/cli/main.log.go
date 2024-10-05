package main

import (
	"os"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar log
	sugar := zap.NewExample().Sugar()
	sugar.Infof("Hello name: %s", "longln")		// same as fmt.Printf

	// logger: key value
	logger := zap.NewExample()
	logger.Info("Hello", zap.String("name", "TipsGo"), zap.Int("age", 40))
	// --> question: when use sugar log, when use logger?

	//  3 types
	// Example

	// Development

	// Production

	// Custom log
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	loggerCustom := zap.New(core, zap.AddCaller())

	loggerCustom.Info("Infor log", zap.Int("line", 1))
	loggerCustom.Error("Error log", zap.Int("line", 2))

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

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./logs/log.txt", os.O_RDWR | os.O_APPEND, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}