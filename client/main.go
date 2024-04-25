package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	// URL of the echo server
	url := "http://10.7.99.33:8663"

	// Get command-line arguments
	args := os.Args[1:]

	// Join arguments into a single string separated by spaces
	data := strings.Join(args, " ")

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return
	}

	// Set the content type header
	req.Header.Set("Content-Type", "text/plain")

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %s\n", err)
		return
	}

	// Print the response body
	fmt.Printf("Response: %s\n", body)
}
