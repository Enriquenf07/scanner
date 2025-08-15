package barcode

import (
	"encoding/json"
	"log"
	"net/http"
	"scanner-backend/utils"

	"github.com/gorilla/mux"
)


func RegisterRoutes(r *mux.Router, middleware func()){
	s := r.PathPrefix("/barcode").Subrouter()
	s.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		middleware()
		ctx := r.Context()
		var data BarcodeRequest
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		log.Println(r.RequestURI)
		err := Create(ctx, data)
		utils.HandleError(err, "erro interno", w)
		defer r.Body.Close()
	}).Methods("POST")

	s.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		middleware()
		ctx := r.Context()
		log.Println(r.RequestURI)
		barcodes, err := GetAll(ctx)
		utils.HandleError(err, "erro interno", w)
		response, err := json.Marshal(barcodes)
		utils.HandleError(err, "erro interno", w)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}).Methods("GET")
}