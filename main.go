package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/regmarmcem/stock-manager/controllers"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbHost     = os.Getenv("DB_HOSTNAME")
	dbPort     = os.Getenv("DB_PORT")
	dbConn     = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbDatabase)
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Println("fail to connect DB")
		return
	}

	fmt.Println(db)

	r := chi.NewRouter()
	r.MethodFunc(http.MethodGet, "/home", controllers.GetStock)

	log.Println("server start at port 8080")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
