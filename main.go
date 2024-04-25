package main

import (
	"fmt"
	"io"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	fmt.Println(string(body))
	// Echo back the request body
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", body)
}

func main() {
	http.HandleFunc("/", echoHandler)
	fmt.Println("Starting server on port 8663...")
	if err := http.ListenAndServe(":8663", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
