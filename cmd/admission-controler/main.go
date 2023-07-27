package main

import (
	"flag"

	"github.com/armosec/admission-controler/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	var tlsKey, tlsCert, port string
	flag.StringVar(&tlsKey, "tlsKey", "/etc/certs/tls.key", "Path to the tls key")
	flag.StringVar(&tlsCert, "tlsCert", "/etc/certs/tls.crt", "Path to the TLS certificate")
	flag.Parse()

	server := server.NewServer(port)
	log.Info().Msgf("Starting server on port: %s", port)
	if err := server.ListenAndServeTLS(tlsCert, tlsKey); err != nil {
		log.Error().Msgf("Failed to listen and serve: %v", err)
		panic(1)
	}
}
