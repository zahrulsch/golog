package golog

import (
	"os"

	"github.com/zahrulsch/golog/encoder"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	*zap.Logger
}

func NewLogger(logLocationPath, appName string) (*Log, error) {
	encoderStdout := zap.NewProductionEncoderConfig()
	encoderStdout.ConsoleSeparator = " "
	encoderStdout.EncodeLevel = encoder.LevelEncoderStdout
	encoderStdout.EncodeTime = encoder.HeadEncoderStdout(appName)

	encoderFile := zap.NewProductionEncoderConfig()
	encoderFile.ConsoleSeparator = " "
	encoderFile.EncodeLevel = encoder.LevelEncoderFile
	encoderFile.EncodeTime = encoder.HeadEncoderFile(appName)

	stdEnc := zapcore.NewConsoleEncoder(encoderStdout)
	fileEnc := zapcore.NewConsoleEncoder(encoderFile)

	fileTarget, err := os.OpenFile(logLocationPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return nil, err
	}

	writer := zapcore.AddSync(fileTarget)
	tees := zapcore.NewTee(
		zapcore.NewCore(fileEnc, writer, zap.DebugLevel),
		zapcore.NewCore(stdEnc, zapcore.AddSync(os.Stdout), zap.DebugLevel),
	)

	jap := zap.New(tees, zap.AddStacktrace(zap.FatalLevel))
	loggerInstance := Log{
		jap,
	}

	return &loggerInstance, nil
}
