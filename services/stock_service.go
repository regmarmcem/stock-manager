package services

import (
	"database/sql"
	"errors"
	"log"

	"github.com/regmarmcem/stock-manager/models"
	"github.com/regmarmcem/stock-manager/repositories"
)

func (s *StockAppService) GetStockService(stockID int) (models.Stock, error) {
	var stock models.Stock
	stock, err := repositories.SelectStocks(s.db, stockID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("Record not found")
			return models.Stock{}, nil
		}
		return models.Stock{}, err
	}
	return stock, nil
}
