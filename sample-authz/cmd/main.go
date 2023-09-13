package main

import (
	"context"
	"flag"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/raanfefu/basajaun/sample-authz/cmd/handler"

	"github.com/common-nighthawk/go-figure"
)

func int64_carg(v *int64) {
	if *v == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	ofigure := figure.NewColorFigure("authz    sample", "o8", "green", true)
	ofigure.Print()

	port := int64(0)

	flag.Int64Var(&port, "port", 0, "port using listing service")
	flag.Parse()
	int64_carg(&port)

	if port == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%v", port),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.HealthCheck)
	mux.HandleFunc("/api/v1/echo", handler.EcoHandler)
	server.Handler = mux

	// Start new go routine for web server
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Printf("Failed to listen and serve webhook server: %v\n", err)
		}
	}()

	fmt.Printf("Server running listening on port: %v\n", port)

	// Listen for shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	// Emit a shutdown log
	fmt.Println("Got shutdown signal, shutting down webhook server gracefully...\n")
	server.Shutdown(context.Background())
}
