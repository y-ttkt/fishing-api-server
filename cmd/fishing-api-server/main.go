package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/yusuke-takatsu/fishing-api-server/config/database"
	"github.com/yusuke-takatsu/fishing-api-server/infra/repository/user"
	"github.com/yusuke-takatsu/fishing-api-server/interface/handler"
	loginUsecase "github.com/yusuke-takatsu/fishing-api-server/usecase/user"
	"io"
	"log"
	"net/http"
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

	userRepo := user.NewRepository(db)
	login := loginUsecase.NewLoginUseCase(userRepo)
	userHandler := handler.NewUserHandler(login)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/v1/login", userHandler.Login)
	if err := http.ListenAndServe(":"+os.Getenv("APP_PORT"), r); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
