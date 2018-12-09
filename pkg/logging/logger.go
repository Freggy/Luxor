package logging

import (
	"github.com/luxordynamics/luxor/bazel-luxor/external/go_sdk/src/strings"
	"go.uber.org/zap"
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

	return nil
}
