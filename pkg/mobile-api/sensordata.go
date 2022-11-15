package mobileapi

import (
	"net/http"
	sensordata "server-poc/pkg/sensor-data"
)

func handleGetAllSensorData(w http.ResponseWriter, r *http.Request) {
	data, err := sensordata.GetAll()
	if err != nil {
		respondError(w, 500, err)
		return
	}
	respondJSON(w, data)
}
