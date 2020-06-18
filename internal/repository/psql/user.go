package psql

import (
	"context"

	storage "github.com/adhistria/auth-movie-app/infrastructure/storage"
	"github.com/adhistria/auth-movie-app/internal/domain"
	log "github.com/sirupsen/logrus"
)

// UserRepository represent user psql
type UserRepository struct {
	DB *storage.Database
}

// Create add new user to database
func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (name, email, password) VALUES(:name, :email, :password);`
	_, err := r.DB.Conn.NamedExec(query, user)

	if err != nil {
		log.Fatalf("Can't create new user: %v", err)
		return err
	}

	return nil
}
