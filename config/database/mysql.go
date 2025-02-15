package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal("Could not initialize database location.")
		return nil, err
	}

	config := mysql.Config{
		DBName:               os.Getenv("DB_DATABASE"),
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Addr:                 os.Getenv("DB_HOST"),
		Net:                  "tcp",
		ParseTime:            true,
		Loc:                  jst,
		AllowNativePasswords: true,
	}

	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
		return nil, err
	}

	return db, nil
}
