package env

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

var Spec Specification

// Specification - env variables struct
type Specification struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
	Port     string `envConfig:"PORT" default:"8080"`
}

// New reads env vars to a struct
func New() *Specification {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Warn("Error loading .env file, will use env vars")
	}

	err = envconfig.Process("", &Spec)
	if err != nil {
		log.Fatal(err)
	}

	return &Spec
}
