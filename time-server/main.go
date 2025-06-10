package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()

	// Format the time in a readable format
	timeString := currentTime.Format("2006-01-02 15:04:05 MST")

	// Set content type to plain text
	w.Header().Set("Content-Type", "text/plain")

	// Write the time response
	fmt.Fprintf(w, "Current time: %s\n", timeString)

	// Log the request
	log.Printf("Request from %s - Time served: %s", r.RemoteAddr, timeString)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Server is healthy!\n")
}

func main() {
	// Set up routes
	http.HandleFunc("/", timeHandler)
	http.HandleFunc("/health", healthHandler)

	port := ":8080"
	log.Printf("Starting server on port %s", port)
	log.Printf("Visit http://localhost%s to get the current time", port)
	log.Printf("Visit http://localhost%s/health for health check", port)

	// Start the server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
