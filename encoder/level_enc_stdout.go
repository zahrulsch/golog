package encoder

import (
	"fmt"

	"github.com/fatih/color"
	"go.uber.org/zap/zapcore"
)

func LevelEncoderStdout(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {
	level := l.CapitalString()
	colorLevel := color.New()

	switch level {
	case "INFO":
		colorLevel.Set().Add(color.FgCyan)
	case "ERROR":
		colorLevel.Set().Add(color.FgRed)
	case "WARN":
		colorLevel.Set().Add(color.FgHiYellow)
	default:
		colorLevel.Set().Add(color.FgWhite)
	}

	result := fmt.Sprintf("[%v]", colorLevel.Sprint(level))
	pae.AppendString(result)
}
