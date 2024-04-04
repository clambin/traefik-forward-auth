package main

import (
	"fmt"
	"net/http"

	internal "github.com/thomseddon/traefik-forward-auth/internal"
)

// Main
func main() {
	// Parse options
	config := internal.NewGlobalConfig()

	// Setup logger
	log := internal.NewDefaultLogger()

	// Perform config validation
	config.Validate()

	// Attach router to default server
	s, err := internal.NewServer()
	if err != nil {
		log.Fatalf("failed to create server: %s", err.Error())
	}
	http.Handle("/", s)

	// Start
	log.WithField("config", config).Debug("Starting with config")
	log.Infof("Listening on :%d", config.Port)
	log.Info(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}
