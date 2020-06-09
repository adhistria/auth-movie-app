package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/adhistria/auth-movie-app/infrastructure/logs"
	"github.com/joho/godotenv"
)

func main() {
	logs.Setup()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Trace(os.Getenv("DATABASE_URL"))
}
