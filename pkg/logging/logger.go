package logging

import "github.com/sirupsen/logrus"

type Config struct {
	LogType        string `json:"logType"`
	LogToJson      bool   `json:"logToJson"`
	PushToLogstash bool   `json:"pushToLogStash"`
}



func newDebugLogger(config Config) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	return logger
}

func newProductionLogger(config Config) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	// TODO: implement the following options:
	// - Use JSON logger, if configured
	// - Use logstash logger, if configured
	return logger
}
