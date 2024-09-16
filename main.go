package main

import (
	"fmt"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"os"
	"tkc/go_bedroxk_proxy_server/server"
)

func main() {
	// Load configuration file
	config, err := server.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create logs directory if it doesn't exist
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		log.Fatalf("Failed to create logs directory: %v", err)
	}

	requestLogger := server.CreateRequestLogger()
	responseLogger := server.CreateResponseLogger()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	// Redirect proxy requests to custom target
	proxy.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		server.RedirectHandler(w, req, config, requestLogger, responseLogger)
	})

	if err := http.ListenAndServe("localhost:"+config.Port, proxy); err != nil {
		fmt.Printf("ListenAndServe: %v", err)
	}
}
