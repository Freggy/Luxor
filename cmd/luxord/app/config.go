package app

import "github.com/luxordynamics/luxor/pkg/logging"

type Config struct {
	HttpsEnabled bool             `json:"httpsEnabled"`
	LogLevel     logging.LogLevel `json:"logLevel"`
}
