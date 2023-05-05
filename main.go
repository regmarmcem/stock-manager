package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/regmarmcem/stock-manager/api"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOSTNAME")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	// TODO: sslmode=disable
	dbConn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbDatabase, dbUser, dbPassword)

	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Println("fail to connect DB")
		return
	}

	r := api.NewRouter(db)
	log.Println("server start at port 8080")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
