package controllers

import (
	"encoding/json"
	"net/http"

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
	stock := &models.Stock{
		ID:   1,
		Name: "stock1",
	}
	json.NewEncoder(w).Encode(stock)
}
