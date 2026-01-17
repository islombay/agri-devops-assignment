package router

import (
	"net/http"

	"github.com/islombay/agri-devops-assignment/internal/handler"
	http2 "github.com/islombay/agri-devops-assignment/internal/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(h *handler.AgricultureHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/farmers", h.Farmers)
	mux.HandleFunc("/api/suppliers", h.Suppliers)
	mux.HandleFunc("/api/products", h.Products)
	mux.HandleFunc("/api/orders", h.Orders)

	mux.Handle("/metrics", promhttp.Handler())

	return http2.PrometheusMiddleware(mux)
}
