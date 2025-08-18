package barcode

import (
	"encoding/json"
	"net/http"
	"scanner-backend/utils"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/barcode", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			barcodes, err := GetAll(ctx)
			utils.HandleError(err, "erro interno", w)
			response, err := json.Marshal(barcodes)
			utils.HandleError(err, "erro interno", w)
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			var data BarcodeRequest
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				http.Error(w, "JSON inválido", http.StatusBadRequest)
				return
			}
			err := Create(ctx, data)
			utils.HandleError(err, "erro interno", w)
			defer r.Body.Close()
		})
	})
}
