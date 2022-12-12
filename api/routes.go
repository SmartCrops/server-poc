package api

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (s *server) routes() {
	s.r.Use(middleware.Logger)

	// Public endpoints
	s.r.Group(func(r chi.Router) {
		s.r.Get("/", s.handleIndex)
	})

	// Auth only endpoints
	s.r.Group(func(r chi.Router) {
		s.r.Get("/users/{userId}/installations", s.handleGetUserInstallations)
	})
}
