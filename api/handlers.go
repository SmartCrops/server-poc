package api

import (
	"net/http"
	"server-poc/pkg/models"
	"strconv"

	"github.com/go-chi/chi"
)

func (s *server) handleGetUserInstallations(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseUint(chi.URLParam(r, "userId"), 10, 64)
	if err != nil {
		s.respondErr(w, err, http.StatusBadRequest)
		return
	}
	user := models.User{}
	err = user.GetByID(s.db, uint(userId))
	if err != nil {
		s.respondErr(w, err, http.StatusInternalServerError)
		return
	}
	s.respondJSON(w, user.Installations)
}

func (s *server) handleIndex(w http.ResponseWriter, r *http.Request) {
	s.respondJSON(w, map[string]string{"message": "Welcome to Smart Crops!"})
}
