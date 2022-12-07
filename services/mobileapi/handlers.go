package mobileapi

import (
	"net/http"
	"server-poc/pkg/models"

	"github.com/go-chi/chi/v5"
)

func (s *server) handleGetSensorData(w http.ResponseWriter, r *http.Request) {
	serial := chi.URLParam(r, "serial")
	data, err := models.GetByDataCollectorSerialNumber(s.db, serial)
	if err != nil {
		s.respondErr(w, err, http.StatusInternalServerError)
		return
	}
	s.respondJSON(w, data)
}

func (s *server) handleIndex(w http.ResponseWriter, r *http.Request) {
	s.respondJSON(w, map[string]string{"message": "Welcome to Smart Crops!"})
}
