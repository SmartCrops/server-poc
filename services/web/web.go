package web

import (
	"fmt"
	"net/http"
	"server-poc/api"
	"server-poc/frontend"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func cors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		handler.ServeHTTP(w, r)
	})
}

func Run(db *gorm.DB, port int) error {
	r := chi.NewRouter()
	r.Mount("/api", api.New(db))
	r.Mount("/", cors(frontend.SvelteKitHandler("/")))
	listenAddr := fmt.Sprintf(":%d", port)
	return http.ListenAndServe(listenAddr, r)
}
