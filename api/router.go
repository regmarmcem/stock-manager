package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/regmarmcem/stock-manager/controllers"
	"github.com/regmarmcem/stock-manager/services"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	ser := services.NewStockAppService(db)
	tCon := controllers.NewStockController(ser)
	r.MethodFunc(http.MethodGet, "/home", tCon.GetStock)
	return r
}
