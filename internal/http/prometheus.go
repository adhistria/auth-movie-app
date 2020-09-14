package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PrometheusHandler represent prometheus handler
type PrometheusHandler struct{}

// Metrics handler
func (p *PrometheusHandler) Metrics(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}

// NewPrometheusHandler return prometheus handler
func NewPrometheusHandler(router *httprouter.Router) {
	ph := PrometheusHandler{}
	router.GET("/metrics", ph.Metrics(promhttp.Handler()))
}
