package services

import (
	"database/sql"

	"github.com/regmarmcem/stock-manager/models"
)

type StockServicer interface {
	GetStockService(stockID int) (models.Stock, error)
}

type StockAppService struct {
	db *sql.DB
}

func NewStockAppService(db *sql.DB) *StockAppService {
	return &StockAppService{db: db}
}
