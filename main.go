package main

import (
	"fmt"
	"github.com/CassioRoos/e-core/handlers"
	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/env"
	"net/http"
	"time"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

var port = env.String("PORT", false, ":8080", "Bind the application to a specific port")

func main() {
	// getting the env vars
	env.Parse()
	// Simple log config
	var log = hclog.New(&hclog.LoggerOptions{
		Name:       "e-core - ",
		Level:      hclog.LevelFromString("DEBUG"),
	})


	server := http.Server{
		Addr:         *port,              // configure the bind address
		Handler:      handlers.ServerMux, // set the default handler
		ReadTimeout:  5 * time.Second,    // max time to read request from the client
		WriteTimeout: 10 * time.Second,   // max time to write response to the client
		IdleTimeout:  120 * time.Second,  // max time for connections using TCP Keep-Alive
	}
	// Showing the user the where is running and how to stop it
	log.Info(fmt.Sprintf("Listening to port %s. Press CTRL + C to stop it.", *port))
	server.ListenAndServe()
}
