package hooks

import (
	"net/http"
	"scanner-backend/monitor"
)

func Monitor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		monitor.HttpRequestsTotal.Inc()
		next.ServeHTTP(w, r)
	})
}
