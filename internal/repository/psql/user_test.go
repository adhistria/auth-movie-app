package psql_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/adhistria/auth-movie-app/infrastructure/storage"
	"github.com/adhistria/auth-movie-app/internal/domain"
	"github.com/adhistria/auth-movie-app/internal/repository/psql"
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
	userRepo := psql.UserRepository{DB: &mockStorage}
	mockUser := domain.User{
		Name:     "adhi satria",
		Email:    "adhistria@gmail.com",
		Password: "password",
	}

	err = userRepo.Create(context.Background(), &mockUser)
	if err != nil {
		t.Errorf("Error was not expected: %s", err)
	}
}
