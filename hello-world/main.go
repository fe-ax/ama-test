package main

import (
	"fmt"
	"io"
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

func timeHandler(w http.ResponseWriter, r *http.Request) {
	// Make request to time-server
	resp, err := http.Get("http://time-server:8080/")
	if err != nil {
		http.Error(w, "Failed to reach time server", http.StatusInternalServerError)
		log.Printf("Error calling time server: %v", err)
		return
	}
	defer resp.Body.Close()

	// Copy the response headers
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Set the status code
	w.WriteHeader(resp.StatusCode)

	// Forward the response body
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Printf("Error forwarding response: %v", err)
		return
	}

	// Log the request
	log.Printf("Request from %s - Time forwarded from time-server", r.RemoteAddr)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Server is healthy!\n")
}

func main() {
	// Set up routes
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/health", healthHandler)

	port := ":8081"
	log.Printf("Starting hello-world server on port %s", port)
	log.Printf("Visit http://localhost%s to get hello world", port)
	log.Printf("Visit http://localhost%s/time to get current time", port)
	log.Printf("Visit http://localhost%s/health for health check", port)

	// Start the server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
