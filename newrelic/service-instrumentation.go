package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/newrelic/go-agent"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}

func main() {
	// Create a config.  You need to provide the desired application name
	// and your New Relic license key.
	cfg := newrelic.NewConfig("helloexample", "4ed3875a608b5d09fcb0289bf9dcba691b9d0f8d")

	// Create an application.  This represents an application in the New
	// Relic UI.
	app, err := newrelic.NewApplication(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Wrap helloHandler.  The performance of this handler will be recorded.
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", helloHandler))
	http.ListenAndServe(":8000", nil)
}
