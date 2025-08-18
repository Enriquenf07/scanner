package main

import (
	"log"
	"net/http"
	"scanner-backend/config"
	"scanner-backend/router"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	config.ConnectRedis()
	router.RegisterRoutes(r)
	log.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", r)
}
