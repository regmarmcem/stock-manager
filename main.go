package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

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
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	// TODO: sslmode=disable
	databaseUrl := fmt.Sprintf("postgresql://%s:%s/%s?user=%s&password=%s&sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPass)

	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Println("fail to connect DB")
		panic(err)
	}

	r := api.NewRouter(db)
	log.Println("server start at port 8080")

	log.Panic(http.ListenAndServe(":8080", r))
}
