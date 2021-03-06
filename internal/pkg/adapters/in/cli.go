package inadapter

import (
	"golang-starter-kit/internal/app/example"
	"log"
)

// CliAdapter - struct with necessary use-cases for adapter to run
type CliAdapter struct {
	example example.Port
}

// NewCliAdapter - create a new instance of NewCliAdapter with passed implementations
func NewCliAdapter(example example.Port) *CliAdapter {
	return &CliAdapter{example: example}
}

// Run - initializes cli adapter run
func (in *CliAdapter) Run() {
	err := in.example.Run()
	if err != nil {
		log.Fatal(err)
	}
}
