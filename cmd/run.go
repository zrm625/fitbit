package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/zrm625/fitbit"
)

const (
	clientSecretKey = "CLIENT_SECRET"
	authCodeKey     = "AUTH_CODE"
	clientIDKey     = "CLIENT_ID"
	databaseURL     = "DATABASE_URL"
)

// TODO: use it
// var db *pgxpool.Pool

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	clientSecret := os.Getenv(clientSecretKey)
	authCode := os.Getenv(authCodeKey)
	clientID := os.Getenv(clientIDKey)
	dbURL := os.Getenv(databaseURL)

	db, err := getConnection(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	f, err := fitbit.New(authCode, clientID, clientSecret)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	w, err := f.GetWeights(start, end)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(w))
}

func getConnection(url string) (*sql.DB, error) {
	return sql.Open("postgres", url)
}
