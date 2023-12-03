package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gogintemplate/env"
	"gopkg.in/natefinch/lumberjack.v2"
)

const Development = "dev"
const Production = "prod"

var logger *zap.Logger

var LogFileWriter = zapcore.AddSync(&lumberjack.Logger{
	Filename:   "logs/gin-service.log",
	MaxSize:    10,
	MaxBackups: 1,
	MaxAge:     1,
})

var LogConsoleWriter = zapcore.AddSync(os.Stdout)

func init() {
	var core zapcore.Core

	switch env.Env.BuildEnv {
	case Development:
		// console and file logs for dev
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			zapcore.NewMultiWriteSyncer(LogConsoleWriter, LogFileWriter),
			zap.InfoLevel,
		)

	case Production:
		// file logs for prod
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
			LogFileWriter,
			zap.InfoLevel,
		)
	}

	logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
}

func GetLogger() *zap.Logger {
	return logger
}
