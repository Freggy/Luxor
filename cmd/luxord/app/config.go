package app

import (
	"encoding/json"
	"github.com/luxordynamics/luxor/pkg/logging"
	"io/ioutil"
)

const ConfigLocation = "/etc/luxor/luxor.json"

type Config struct {
	HttpsEnabled bool           `json:"httpsEnabled"`
	LoggerConf   logging.Config `json:"loggerConfiguration"`
}


func NewDefaultConfig() *Config {
	return nil
}

// ConfigFromFile reads the the file at the given path. The content needs to be in JSON format.
func ConfigFromFile(path string) (*Config, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	conf, err := fromString(string(content))

	if err != nil {
		return nil, err
	}

	return conf, nil
}

func fromString(content string) (*Config, error) {
	conf := Config{}

	if err := json.Unmarshal([]byte(content), &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
