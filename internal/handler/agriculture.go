package handler

import (
	"encoding/json"
	"net/http"

	"github.com/islombay/agri-devops-assignment/internal/service"
)

type AgricultureHandler struct {
	service *service.AgricultureService
}

func NewAgricultureHandler(s *service.AgricultureService) *AgricultureHandler {
	return &AgricultureHandler{service: s}
}

func (h *AgricultureHandler) Farmers(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(h.service.GetFarmers())
}

func (h *AgricultureHandler) Suppliers(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(h.service.GetSuppliers())
}

func (h *AgricultureHandler) Products(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(h.service.GetProducts())
}

func (h *AgricultureHandler) Orders(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(h.service.GetOrders())
}
