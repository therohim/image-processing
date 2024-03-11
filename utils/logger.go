package utils

import (
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func LogAny(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func ByteString(key string, value []byte) zap.Field {
	return zap.ByteString(key, value)
}

func LogString(key string, value string) zap.Field {
	return zap.String(key, value)
}

func NewLogger() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := ecszap.NewDefaultEncoderConfig()
	writer := zapcore.AddSync(os.Stdout)
	defaultLogLevel := zapcore.DebugLevel
	core := ecszap.NewCore(encoder, writer, defaultLogLevel)

	// DON'T REMOVE THIS CODE!!!
	// core := zapcore.NewTee(
	// 	zapcore.NewCore(jsonEncoder, writer, defaultLogLevel),
	// )

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
}
