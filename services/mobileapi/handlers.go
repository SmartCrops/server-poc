package mobileapi

import (
	"net/http"
	"server-poc/pkg/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s *server) handleGetSensorData(w http.ResponseWriter, r *http.Request) {
	sensorID, err := strconv.ParseInt(chi.URLParam(r, "sensorID"), 10, 64)
	if err != nil {
		s.respondErr(w, err, http.StatusBadRequest)
		return
	}
	data, err := models.GetBySensorID(s.db, int(sensorID))
	if err != nil {
		s.respondErr(w, err, http.StatusInternalServerError)
		return
	}
	s.respondJSON(w, data)
}

func (s *server) handleIndex(w http.ResponseWriter, r *http.Request) {
	s.respondJSON(w, map[string]string{"message": "Welcome to Smart Crops!"})
}