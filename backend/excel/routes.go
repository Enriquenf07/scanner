package excel

import (
	"net/http"
	"scanner-backend/utils"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, middleware func()){
	s := r.PathPrefix("/excel").Subrouter()
	s.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		middleware()
		ctx := r.Context()
		err := Create(ctx)
		utils.HandleError(err, "erro interno", w)
	})
}