package main

import "github.com/luxordynamics/luxor/pkg/logging"

func main() {

	logger := logging.NewProductionLogger()

	logger.Info("Hello World!")
	logger.Debug("DEBUG!!!")
	logger.Warn("WARNING!!")
	logger.Error("HELLO")
}
