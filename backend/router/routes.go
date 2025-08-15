package router

import (
	"scanner-backend/barcode"
	"scanner-backend/monitor"

	"github.com/gorilla/mux"
)

func middlewares(){
	monitor.HttpRequestsTotal.Inc()
}

func RegisterRoutes(r *mux.Router) {
    barcode.RegisterRoutes(r, middlewares)
	monitor.RegisterRoutes(r)
}