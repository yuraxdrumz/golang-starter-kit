package httpadapter

import (
	"github.com/yuraxdrumz/golang-starter-kit/internal/app/example"
)

// NewCliAdapter - create a new instance of NewCliAdapter with passed implementations
func NewCliAdapter(example *example.Something) *CliAdapter {
	return &CliAdapter{example: *example}
}

// Run - initializes cli adapter run
func (in *CliAdapter) Run() {
	in.example.Run()
}
