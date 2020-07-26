package psql

import (
	"context"

	storage "github.com/adhistria/auth-movie-app/infrastructure/storage"
	"github.com/adhistria/auth-movie-app/internal/domain"
	log "github.com/sirupsen/logrus"
)

// UserRepository represent user psql
type userRepository struct {
	DB *storage.Database
}

// Create add new user to database
func (r userRepository) Create(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (name, email, password) VALUES(:name, :email, :password);`
	_, err := r.DB.Conn.NamedExec(query, user)

	if err != nil {
		log.Warnf("Can't create new user: %v", err)
		return err
	}

	return nil
}

// FindByEmail find user by email
func (r userRepository) FindByEmail(ctx context.Context, user *domain.User) (*domain.User, error) {
	newUser := domain.User{}
	query := `SELECT email, password FROM users WHERE email = ?;`
	err := r.DB.Conn.Get(&newUser, query, user.Email)

	if err != nil {
		log.Warnf("Can't find user with email: %v", user.Email)
		return nil, err
	}

	return &newUser, nil
}

// NewUserRepository ...
func NewUserRepository(db *storage.Database) domain.UserRepository {
	return &userRepository{DB: db}
}
