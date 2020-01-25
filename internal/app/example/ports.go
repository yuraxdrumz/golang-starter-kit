package example

import "time"

// FileExists - check if file exists type
type FileExists func(filename string) bool

// Sleep - sleep for duration
type Sleep func(time time.Duration)

// Something - struct with neccessary implementations to return new methods
type Something struct {
	checkFileExists FileExists
	sleeper         Sleep
}
