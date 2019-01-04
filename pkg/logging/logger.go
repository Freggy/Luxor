package logging

import "github.com/sirupsen/logrus"

type Config struct {
	LogType        string `json:"logType"`
	LogToJson      bool   `json:"logToJson"`
	PushToLogstash bool   `json:"pushToLogStash"`
}


func NewLoggerFromType(logType string, config *Config) *logrus.Logger {
	if logType == "DEBUG" {
		return NewDebugLogger()
	} else if logType == "PRODUCTION" {
		return NewProductionLogger(config)
	} else {
		return NewDebugLogger()
	}
}

func NewDebugLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	return logger
}

func NewProductionLogger(config *Config) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	// TODO: implement the following options:
	// - Use JSON logger, if configured
	// - Use logstash logger, if configured
	return logger
}
