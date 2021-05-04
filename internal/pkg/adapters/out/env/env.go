package env

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var spec Specification

// Specification - env variables struct
type Specification struct {
	LogzioToken string `envconfig:"LOGZIO_TOKEN"`
	AppName     string `envconfig:"APP_NAME" default:"example-app"`
	LogLevel    string `envconfig:"LOG_LEVEL" default:"info"`
	Port        string `envConfig:"PORT" default:"8080"`
}

// GetEnv returns local spec
func (s *Specification) GetEnv() *Specification {
	return &spec
}


func New() *Specification {
	err := envconfig.Process("", &spec)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return &spec
}