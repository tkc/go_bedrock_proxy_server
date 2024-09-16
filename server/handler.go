package server

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// RedirectHandler handles the proxying of requests and logging of both the request and response.
func RedirectHandler(w http.ResponseWriter, r *http.Request, conf *Config, requestLogger *RequestLogger, responseLogger *ResponseLogger) {
	// Log the request
	requestLogger.Printf(r)

	// Read the request body
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("error reading request body: %w", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Sign the request
	req, err := sign(bodyBytes, conf)
	if err != nil {
		fmt.Errorf("error signing request: %w", err)
		http.Error(w, "Failed to sign request", http.StatusInternalServerError)
		return
	}

	// Create an HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second, // Set timeout to avoid long requests
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("error sending request: %w", err)
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		responseLogger.Printf("Error reading response body: %v", err)
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	// Log the response
	responseLogger.Printf("Response Status: %d, Body: %s", resp.StatusCode, string(respBodyBytes))

	// Copy the response headers to the client
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Set the response status code
	w.WriteHeader(resp.StatusCode)

	// Write the response body to the client
	if _, err := w.Write(respBodyBytes); err != nil {
		responseLogger.Printf("Error writing response body to client: %v", err)
	}
}
