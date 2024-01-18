package main

import (
	"context"
	"flag"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/common-nighthawk/go-figure"
	"github.com/gorilla/mux"
	"github.com/raanfefu/basajaun/gui/cmd/handler"
)

func int64_carg(v *int64) {
	if *v == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	ofigure := figure.NewColorFigure("authz template", "o8", "green", true)
	ofigure.Print()

	port := int64(0)
	directory := flag.String("directory", "/private/static/", "the directory of static file to host")
	flag.Int64Var(&port, "port", 0, "port using listing service")
	flag.Parse()

	int64_carg(&port)

	fmt.Println(*directory)

	mux := mux.NewRouter()

	mux.HandleFunc("/api/v1/publish", handler.PostPublish).Methods("POST")
	mux.HandleFunc("/api/v1/policy", handler.PostPolicy).Methods("POST")
	mux.HandleFunc("/api/v1/settings", handler.GetSettings).Methods("GET")

	mux.PathPrefix("/").HandlerFunc(handler.EnableCors).Methods("OPTIONS")
	mux.HandleFunc("/api/v1/health", handler.HealthCheck)
	mux.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(*directory))))

	server := &http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%v", port),
	}

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
