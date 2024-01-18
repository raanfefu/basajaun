package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/raanfefu/basajaun/admission/cmd/handler"

	"github.com/common-nighthawk/go-figure"
)

func string_carg(v *string) {
	if *v == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func int64_carg(v *int64) {
	if *v == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	ofigure := figure.NewColorFigure("authz    opa", "o8", "green", true)
	ofigure.Print()
	crt := ""
	key := ""
	port := int64(0)

	flag.StringVar(&crt, "crt", "", "the cert file path")
	flag.StringVar(&key, "key", "", "The key file")
	flag.Int64Var(&port, "port", 0, "port using listing service")
	flag.Parse()
	string_carg(&crt)
	string_carg(&key)
	int64_carg(&port)

	if key == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if port == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Load these into a certs object
	certs, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		fmt.Printf("Failed to load key pair: %v\n", err)
	}

	server := &http.Server{
		Addr:      fmt.Sprintf(":%v", port),
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{certs}},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.HealthCheck)
	mux.HandleFunc("/mutation", handler.Mutation)
	server.Handler = mux

	// Start new go routine for web server
	go func() {
		if err := server.ListenAndServeTLS("", ""); err != nil {
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
