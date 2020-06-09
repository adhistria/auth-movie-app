package storage

import (
	"os"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

// Database struct
type Database struct {
	Conn *sqlx.DB
}

// NewConnection return database connection
func NewConnection() (*Database, error) {
	var db Database
	data, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	db.Conn = data
	if err != nil {
		log.Fatal(err)
	}
	return &db, err
}
