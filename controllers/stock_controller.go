package controllers

import (
	"encoding/json"
	"net/http"
)

type Stock struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	stock := &Stock{
		Name:        "stock1",
		Description: "description",
	}
	json.NewEncoder(w).Encode(stock)
}
