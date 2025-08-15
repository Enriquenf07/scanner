package router

import (
	"net/http"
	"scanner-backend/barcode"
	"scanner-backend/excel"
	"scanner-backend/monitor"

	"github.com/gorilla/mux"
)

func middlewares(){
	monitor.HttpRequestsTotal.Inc()
}

func RegisterRoutes(r *mux.Router) {
    barcode.RegisterRoutes(r, middlewares)
	monitor.RegisterRoutes(r)
	excel.RegisterRoutes(r, middlewares)
	fs := http.FileServer(http.Dir("./data"))
	r.PathPrefix("/data/").Handler(http.StripPrefix("/data/", fs))
}