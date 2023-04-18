package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/regmarmcem/stock-manager/controllers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	r := chi.NewRouter()
	r.MethodFunc(http.MethodGet, "/home", controllers.GetStock)

	log.Println("server start at port 8080")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
