package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to plain text
	w.Header().Set("Content-Type", "text/plain")

	// Write the hello world response
	fmt.Fprintf(w, "hello world\n")

	// Log the request
	log.Printf("Request from %s - Hello world served", r.RemoteAddr)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Server is healthy!\n")
}

func main() {
	// Set up routes
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)

	port := ":8081"
	log.Printf("Starting hello-world server on port %s", port)
	log.Printf("Visit http://localhost%s to get hello world", port)
	log.Printf("Visit http://localhost%s/health for health check", port)

	// Start the server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
