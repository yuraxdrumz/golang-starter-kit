package main

import (
	"github.com/bshuster-repo/logruzio"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/yuraxdrumz/golang-starter-kit/internal/app/example"
	httpadapter "github.com/yuraxdrumz/golang-starter-kit/internal/pkg/adapters/in"
	fileutils "github.com/yuraxdrumz/golang-starter-kit/internal/pkg/adapters/out/fileutils"
	sleeper "github.com/yuraxdrumz/golang-starter-kit/internal/pkg/adapters/out/sleeper"
	"os"
)

// Specification - env variables struct
type Specification struct {
	LogzioToken string `envconfig:"LOGZIO_TOKEN"`
	AppName     string `envconfig:"APP_NAME" default:"example-app"`
	LogLevel    string `envconfig:"LOG_LEVEL" default:"info"`
	Port        string `envConfig:"PORT" default:"8080"`
}

var s Specification

func init() {

	err := envconfig.Process("", &s)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
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

	contextLogger := log.Fields{
		"app_name": s.AppName,
	}

	if s.LogzioToken != "" {
		log.Info("Adding logz.io token")
		hook, err := logruzio.New(s.LogzioToken, s.AppName, contextLogger)
		if err != nil {
			log.Fatal(err)
		}
		log.AddHook(hook)
	}
}

// // with http adapter
// func main() {
// 	log.Debug("init use cases")
// 	// init fileutils
// 	fu := fileutils.NewFileUtilsAdapter()
// 	// init sleeper
// 	sl := sleeper.NewSleepAdapter()
// 	// init scanner use case with provided adapter functions
// 	ex := example.NewExample(fu.FileExists, sl.Sleep)
// 	log.Debug("init in adapters")
// 	// init in adapter with all use cases
// 	inadapter := httpadapter.NewHTTPAdapter(ex)
// 	// run
// 	log.Debug("run in adapters")
// 	inadapter.Run(s.Port)
// }

// with cli adapter
func main() {
	log.Debug("init use cases")
	// init fileutils
	fu := fileutils.NewFileUtilsAdapter()
	// init sleeper
	sl := sleeper.NewSleepAdapter()
	// init example use case with provided adapter functions
	ex := example.NewExample(fu.FileExists, sl.Sleep)
	log.Debug("init in adapters")
	// init in adapter with all use cases
	inadapter := httpadapter.NewCliAdapter(ex)
	// run
	log.Debug("run in adapters")
	inadapter.Run()
}
