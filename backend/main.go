package main

import (
	"encoding/json"
	"log"
	"net/http"
	"scanner-backend/barcode"
	"scanner-backend/config"
	"scanner-backend/utils"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	config.ConnectRedis()

	r.HandleFunc("/barcode", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var data barcode.BarcodeRequest
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		log.Println(r.RequestURI)
		err := barcode.Create(ctx, data)
		utils.HandleError(err, "erro interno", w)
		defer r.Body.Close()
	}).Methods("POST")

	r.HandleFunc("/barcode", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log.Println(r.RequestURI)
		barcodes, err := barcode.GetAll(ctx)
		utils.HandleError(err, "erro interno", w)
		response, err := json.Marshal(barcodes)
		utils.HandleError(err, "erro interno", w)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}).Methods("GET")

	log.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", r)
}
