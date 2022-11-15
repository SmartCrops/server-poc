package mobileapi

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const addr = ":8080"

// Warning - this function will block
func Run() error {
	r := chi.NewRouter()
	r.Get("/sensordata", handleGetAllSensorData)
	log.Println("mobile-api: binding on address", addr)
	return http.ListenAndServe(addr, r)
}
