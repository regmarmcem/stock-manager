package repositories

import (
	"database/sql"
	"log"

	"github.com/regmarmcem/stock-manager/models"
)

func SelectStocks(db *sql.DB, stockID int) (models.Stock, error) {
	const sqlStr = `
		select id, name
		from stocks
		wheree stock_id = ?;
	`
	row := db.QueryRow(sqlStr, stockID)
	if err := row.Err(); err != nil {
		log.Println("SelectStocks failed")
		return models.Stock{}, err
	}
	var stock models.Stock
	err := row.Scan(&stock.ID, &stock.Name)
	if err != nil {
		return models.Stock{}, err
	}

	return stock, nil
}
