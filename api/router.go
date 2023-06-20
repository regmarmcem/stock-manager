package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/regmarmcem/stock-manager/controllers"
	"github.com/regmarmcem/stock-manager/middlewares"
	"github.com/regmarmcem/stock-manager/services"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middlewares.LoggingMiddleware)
	ser := services.NewStockAppService(db)
	tCon := controllers.NewStockController(ser)
	r.MethodFunc(http.MethodGet, "/stock/{id:[0-9]+}", tCon.GetStock)
	return r
}
