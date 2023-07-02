package encoder

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

func LevelEncoderFile(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {
	level := l.CapitalString()
	pae.AppendString(fmt.Sprintf("%s", level))
}
