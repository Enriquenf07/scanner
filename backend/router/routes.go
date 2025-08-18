package router

import (
	"net/http"
	"scanner-backend/barcode"
	"scanner-backend/excel"
	"scanner-backend/hooks"
	"scanner-backend/monitor"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRoutes(r *chi.Mux) {
	r.Route("/metrics", func(r chi.Router) {
		monitor.RegisterRoutes(r)
	})

	r.Route("/", func(r chi.Router) {
		r.Use(hooks.Monitor)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(middleware.RequestID)
		r.Use(middleware.RealIP)
		r.Use(middleware.Timeout(5 * time.Second))
		r.Use(middleware.StripSlashes)

		barcode.RegisterRoutes(r)
		excel.RegisterRoutes(r)

		fs := http.FileServer(http.Dir("./data"))
		r.Handle("/data/*", http.StripPrefix("/data/", fs))
	})
}
