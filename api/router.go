package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/regmarmcem/stock-manager/controllers"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	r.MethodFunc(http.MethodGet, "/home", controllers.GetStock)
	return r
}
