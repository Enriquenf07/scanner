package excel

import (
	"net/http"
	"scanner-backend/utils"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/excel", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			err := Create(ctx)
			utils.HandleError(err, "erro interno", w)
		})
	})
}
