package services

import "github.com/regmarmcem/stock-manager/models"

func (*StockAppService) GetStockService(stockID int) (models.Stock, error) {
	return models.Stock{}, nil
}
