package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/regmarmcem/stock-manager/services"
)

type StockController struct {
	service services.StockServicer
}

func NewStockController(s services.StockServicer) *StockController {
	return &StockController{service: s}
}

func (s *StockController) GetStock(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Path parameter must be a number", http.StatusBadRequest)
		return
	}

	stock, err := s.service.GetStockService(id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(stock)
}
