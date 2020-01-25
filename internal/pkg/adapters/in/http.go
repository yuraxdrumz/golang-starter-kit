package httpadapter

import (
	"net/http"
	"strings"

	"github.com/yuraxdrumz/golang-starter-kit/internal/app/example"
	log "github.com/sirupsen/logrus"
)

// NewHTTPAdapter - create a new instance of NewHTTPAdapter with passed implementations
func NewHTTPAdapter(example *example.Something) *HTTPAdapter {
	return &HTTPAdapter{example: *example}
}

// Run - initializes http server with port
func (in *HTTPAdapter) Run(port string) {
	sayHello := func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Path
		// call example use case passed to in adapter
		in.example.Run()
		message = strings.TrimPrefix(message, "/")
		message = "Hello " + message

		w.Write([]byte(message))
	}
	http.HandleFunc("/", sayHello)
	log.Info("Starting server on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
