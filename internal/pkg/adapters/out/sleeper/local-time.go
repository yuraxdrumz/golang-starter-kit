package sleeper

import (
	"time"
)

// NewSleepAdapter - create a new instance of FileUtilsAdapter with passed implementations
func NewSleepAdapter() *LocalTime {
	return &LocalTime{}
}

// Sleep - sleeps for X duration
func (lt *LocalTime) Sleep(sec time.Duration) {
	time.Sleep(sec)
}
