package server

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

// RequestLogger logs HTTP request details to a log file and stdout.
type RequestLogger struct {
	logger *log.Logger
}

// CreateRequestLogger initializes a RequestLogger to log request details to a file and stdout.
func CreateRequestLogger() *RequestLogger {
	file, err := os.OpenFile("logs/request.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open request log file: %v", err)
	}
	return &RequestLogger{
		logger: log.New(io.MultiWriter(os.Stdout, file), "REQUEST: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Printf logs the HTTP request method, URL, headers, and body.
func (l *RequestLogger) Printf(r *http.Request) {
	l.logger.Printf("Request: %s %s", r.Method, r.URL)
	for name, values := range r.Header {
		for _, value := range values {
			l.logger.Printf("Request Header: %s: %s", name, value)
		}
	}

	// Read and log the request body
	body, err := io.ReadAll(r.Body) // ioutil.ReadAll is deprecated, so io.ReadAll is preferred
	if err != nil {
		l.logger.Printf("Failed to read request body: %v", err)
		return
	}

	// Restore the request body for further processing
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	l.logger.Printf("Request Body: %s", body)
}

// ResponseLogger logs HTTP response details to a log file and stdout.
type ResponseLogger struct {
	logger *log.Logger
}

// CreateResponseLogger initializes a ResponseLogger to log response details to a file and stdout.
func CreateResponseLogger() *ResponseLogger {
	file, err := os.OpenFile("logs/response.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open response log file: %v", err)
	}
	return &ResponseLogger{
		logger: log.New(io.MultiWriter(os.Stdout, file), "RESPONSE: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Printf logs the formatted response details.
func (l *ResponseLogger) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}
