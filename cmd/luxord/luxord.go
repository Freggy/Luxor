package main

import (
	"encoding/json"
	"github.com/luxordynamics/luxor/cmd/luxord/app"
	"github.com/luxordynamics/luxor/pkg/logging"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var config *app.Config
	var logger *logrus.Logger

	if _, err := os.Stat(app.ConfigLocation); os.IsNotExist(err) {
		config = app.NewDefaultConfig()

		data, err := json.Marshal(config)

		if err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile(app.ConfigLocation, data, 777); err != nil {
			log.Fatal(err)
		}
	} else {
		config, err = app.ConfigFromFile(app.ConfigLocation)

		if err != nil {
			log.Fatal(err)
		}
	}

	logger = logging.NewLoggerFromType(config.LoggerConf.LogType, config.LoggerConf)
}
