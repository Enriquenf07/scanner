package monitor

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	HttpRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Número total de requisições HTTP recebidas",
		},
	)
)

func init() {
	prometheus.MustRegister(HttpRequestsTotal)
}

func RegisterRoutes(r *mux.Router) {
	r.Handle("/metrics", promhttp.Handler()).Methods("GET")
}
