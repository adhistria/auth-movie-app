package user_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/adhistria/auth-movie-app/infrastructure/storage"
	"github.com/adhistria/auth-movie-app/internal/user"
	"github.com/jmoiron/sqlx"
)

func TestRegisterUser(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error %s", err)
	}

	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mockDB.ExpectExec("INSERT INTO users").WithArgs("adhi satria", "adhistria@gmail.com", "password").WillReturnResult(sqlmock.NewResult(1, 1))

	mockStorage := storage.Database{Conn: sqlxDB}
	userRepo := user.Repository{DB: &mockStorage}

	mockUser := user.User{
		Name:     "adhi satria",
		Email:    "adhistria@gmail.com",
		Password: "password",
	}

	_, err = userRepo.CreateUser(context.Background(), &mockUser)
	if err != nil {
		t.Errorf("Error was not expected: %s", err)
	}
}
