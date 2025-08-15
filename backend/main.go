package main

import (
	"log"
	"net/http"
	"os"
	"scanner-backend/config"
	"scanner-backend/router"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	config.ConnectRedis()
	router.RegisterRoutes(r)
	log.Println("REDIS_ADDRESS =", os.Getenv("REDIS_ADDRESS"))
	log.Println("REDIS_PASSWORD =", os.Getenv("REDIS_PASSWORD"))
	log.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", r)
}
