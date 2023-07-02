package encoder

import (
	"fmt"
	"time"

	"go.uber.org/zap/zapcore"
)

func HeadEncoderFile(appName string) func(time.Time, zapcore.PrimitiveArrayEncoder) {
	app_name := fmt.Sprintf("%s:", appName)

	return func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		format_date := "02-Jan-2006 15:04:05"
		pae.AppendString(fmt.Sprintf("%s %s", app_name, t.Format(format_date)))
	}
}
