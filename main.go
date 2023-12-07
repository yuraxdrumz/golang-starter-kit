package main

import (
	"os"

	"golang-starter-kit/internal/app/example"
	inadapter "golang-starter-kit/internal/pkg/adapters/in"
	"golang-starter-kit/internal/pkg/adapters/out/env"
	"golang-starter-kit/internal/pkg/adapters/out/fileutils"
	"golang-starter-kit/internal/pkg/adapters/out/sleeper"

	log "github.com/sirupsen/logrus"
)

func init() {
	s := env.New()
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	switch s.LogLevel {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	}
}

// with http adapter
func main() {
	// declare all ports
	var fu fileutils.Port
	var sl sleeper.Port
	var ex example.Port
	var ia inadapter.Port

	log.Debug("init use cases")
	// init fileutils
	fu = fileutils.NewFileUtilsAdapter()
	// init sleeper
	sl = sleeper.NewSleepAdapter()
	// init example use case with provided adapters
	ex = example.NewExample(fu, sl)
	log.Debug("init in adapters")
	// http adapter
	ia = inadapter.NewHTTPAdapter(ex, env.Spec.Port)
	// cli adapter
	//ia = inadapter.NewCliAdapter(ex)
	// run
	log.Debug("run in adapters")
	ia.Run()
}
