package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type server struct {
	db *gorm.DB
	r  chi.Router
}

func New(db *gorm.DB) http.Handler {
	s := &server{
		db: db,
		r:  chi.NewRouter(),
	}
	s.routes()
	return s
}

func (s *server) respondJSON(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		s.respondErr(w, err, http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(b)
}

func (s *server) respondErr(w http.ResponseWriter, err error, code int) {
	type Response struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
	}
	r := Response{code, err.Error()}
	s.respondJSON(w, r)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
