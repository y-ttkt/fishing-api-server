package main

import (
	"github.com/joho/godotenv"
	"github.com/yusuke-takatsu/fishing-api-server/config/database"
	"io"
	"log"
	"os"
)

func init() {
	loadEnv()
	if os.Getenv("APP_ENV") == "production" {
		return
	}

	f, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(io.MultiWriter(f, os.Stdout))
}

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
