package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/douglasmakey/oauth2-example/handlers"
)

// main is the entry point of the application.
func main() {
	// We create a simple server using http.Server and run.
	server := &http.Server{
		Addr:    fmt.Sprintf(":8000"),
		Handler: handlers.New(),
	}

	// Log the start of the server and listen for connections.
	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
