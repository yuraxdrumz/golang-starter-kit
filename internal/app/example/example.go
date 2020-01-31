package example

import (
	log "github.com/sirupsen/logrus"
	"time"
)

// NewExample - create a new instance of Example with passed implementations
func NewExample(checkFileExists FileExists, sleeper Sleep) *Something {
	return &Something{checkFileExists: checkFileExists, sleeper: sleeper}
}

// Run - add method on provided example reference
func (ex *Something) Run() error {
	log.Debug("Running in Example.Run")
	ex.checkFileExists("random")
	ex.sleeper(time.Duration(10))
	return nil
}
