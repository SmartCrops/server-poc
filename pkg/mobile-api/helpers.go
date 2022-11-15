package mobileapi

import (
	"encoding/json"
	"net/http"
)

func respondError(w http.ResponseWriter, code int, err error) {
	type response struct {
		Error string `json:"error"`
	}
	b, _ := json.Marshal(response{err.Error()})
	w.WriteHeader(code)
	w.Write(b)
}

func respondJSON(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		respondError(w, 500, err)
		return
	}
	w.Write(b)
}
