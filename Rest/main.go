package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/JuanObandoDeveloper/rest/handlers"
	"github.com/JuanObandoDeveloper/rest/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_KEY := os.Getenv("JWT_KEY")
	DB_URL := os.Getenv("DB_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTKey: JWT_KEY,
		Port:   PORT,
		DBUrl:  DB_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler()).Methods(http.MethodGet)
}
