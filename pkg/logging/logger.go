package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"os"
	"strings"
)

type LogLevel string

const (
	Debug      LogLevel = "Debug"
	Production LogLevel = "Production"
)

func LogLevelFromString(level string) LogLevel {
	if strings.EqualFold("Debug", level) {
		return Debug
	} else if strings.EqualFold("Production", level) {
		return Production
	}
	return Debug
}

func NewProductionLogger() *zap.Logger {

	// TODO: use custom logger config

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})

	warnLevelAndAbove := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	jsonInfo := zapcore.AddSync(ioutil.Discard)
	jsonErrors := zapcore.AddSync(ioutil.Discard)

	consoleInfo := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, jsonErrors, warnLevelAndAbove),
		zapcore.NewCore(jsonEncoder, jsonInfo, infoLevel),

		zapcore.NewCore(consoleEncoder, consoleErrors, warnLevelAndAbove),
		zapcore.NewCore(consoleEncoder, consoleInfo, infoLevel),
	)

	return zap.New(core)
}
