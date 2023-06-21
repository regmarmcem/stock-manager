package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/regmarmcem/stock-manager/controllers"
	"github.com/regmarmcem/stock-manager/middlewares"
	"github.com/regmarmcem/stock-manager/services"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middlewares.LoggingMiddleware)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))
	ser := services.NewStockAppService(db)
	tCon := controllers.NewStockController(ser)
	r.MethodFunc(http.MethodGet, "/stock/{id:[0-9]+}", tCon.GetStock)
	return r
}
