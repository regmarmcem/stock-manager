package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/regmarmcem/stock-manager/models"
)

func SelectStocks(db *sql.DB, stockID int) (models.Stock, error) {
	const sqlStr = `
		select id, name
		from stocks
		where id = $1;
	`
	row := db.QueryRow(sqlStr, stockID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
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
