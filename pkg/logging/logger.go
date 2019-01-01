package logging

import "github.com/sirupsen/logrus"

func newProductionLogger() *logrus.Logger {
	logger := logrus.New()

	return logger
}
