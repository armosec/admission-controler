package server

import (
	"fmt"
	"net/http"

	"github.com/armosec/admission-controller/internal/server/handlers"
)

func NewServer(port string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/armo-admission-controller/validating", handlers.AdmissionControllerHandler)
	mux.HandleFunc("/armo-admission-controller/mutating", handlers.AdmissionControllerHandler)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
}
