package golog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Any(key string, value interface{}) zapcore.Field {
	return zap.Any(key, value)
}

func String(key, value string) zapcore.Field {
	return zap.String(key, value)
}

func Int(key string, value int) zapcore.Field {
	return zap.Int(key, value)
}

func Bool(key string, value bool) zapcore.Field {
	return zap.Bool(key, value)
}

func Float64(key string, value float64) zapcore.Field {
	return zap.Float64(key, value)
}

func Float32(key string, value float32) zapcore.Field {
	return zap.Float32(key, value)
}

func ByteString(key string, value []byte) zapcore.Field {
	return zap.ByteString(key, value)
}
