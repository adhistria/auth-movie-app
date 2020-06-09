package user

import (
	"context"

	log "github.com/sirupsen/logrus"

	storage "github.com/adhistria/auth-movie-app/infrastructure/storage"
)

// Repository represent user sql
type Repository struct {
	DB *storage.Database
}

// CreateUser add new user to database
func (r *Repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	query := `INSERT INTO users (name, email, password) VALUES(:name, :email, :password) RETURNING id;`
	res, err := r.DB.Conn.NamedExec(query, user)
	if err != nil {
		log.Fatalf("Can't create new user: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("Can't get last user id: %v", err)
	}

	user.ID = uint64(id)
	return user, nil
}
