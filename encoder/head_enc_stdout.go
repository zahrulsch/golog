package encoder

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"go.uber.org/zap/zapcore"
)

func HeadEncoderStdout(appName string) func(time.Time, zapcore.PrimitiveArrayEncoder) {
	magenta := color.New(color.FgHiMagenta)
	app_name := magenta.Sprintf("%s:", appName)

	return func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		format_date := "02-Jan-2006 15:04:05"
		pae.AppendString(fmt.Sprintf("%s %s", app_name, t.Format(format_date)))
	}
}
