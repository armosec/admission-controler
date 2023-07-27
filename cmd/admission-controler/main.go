package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/armosec/admission-controler/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	var tlsKey, tlsCert, port string
	flag.StringVar(&tlsKey, "tlsKey", "/etc/certs/tls.key", "Path to the tls key")
	flag.StringVar(&tlsCert, "tlsCert", "/etc/certs/tls.crt", "Path to the TLS certificate")
	flag.StringVar(&port, "port", "8443", "The port on which to listen")
	flag.Parse()

	server := server.NewServer(port)

	go func() {
		// listen shutdown signal
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		sig := <-signalChan
		log.Error().Msgf("Received %s signal; shutting down...", sig)
		if err := server.Shutdown(context.Background()); err != nil {
			log.Err(err)
		}
	}()

	log.Info().Msgf("Starting server on port: %s", port)
	if err := server.ListenAndServeTLS(tlsCert, tlsKey); err != nil {
		log.Error().Msgf("Failed to listen and serve: %v", err)
		panic(1)
	}
}
