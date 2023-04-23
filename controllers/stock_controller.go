package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/regmarmcem/stock-manager/models"
	"github.com/regmarmcem/stock-manager/services"
)

type StockController struct {
	service services.StockServicer
}

func NewStockController(s services.StockServicer) *StockController {
	return &StockController{service: s}
}

func (*StockController) GetStock(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Path parameter must be a number", http.StatusBadRequest)
		return
	}
	stock := &models.Stock{
		ID:   id,
		Name: "stock1",
	}
	json.NewEncoder(w).Encode(stock)
}
