package main

import (
	"os"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"

	"net/http"

	"github.com/adhistria/auth-movie-app/infrastructure/logs"
	"github.com/adhistria/auth-movie-app/infrastructure/storage"
	http_app "github.com/adhistria/auth-movie-app/internal/http"
	"github.com/adhistria/auth-movie-app/internal/repository/psql"
	"github.com/adhistria/auth-movie-app/internal/service"
	"github.com/adhistria/auth-movie-app/internal/validation"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	logs.Setup()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Trace(os.Getenv("DATABASE_URL"))

	database, err := storage.NewConnection()
	if err != nil {
		log.Fatalf("Error init database connection : %s ", err)
	}
	router := httprouter.New()
	userRepo := psql.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	validator := validation.NewValidator()
	http_app.NewUserHandler(router, userService, validator)
	http_app.NewPrometheusHandler(router)
	log.Info("Run application")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
