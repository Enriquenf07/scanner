package monitor

import (
	"net/http"

	"github.com/go-chi/chi/v5"
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

func RegisterRoutes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		promhttp.Handler().ServeHTTP(w, r)
	})
}
