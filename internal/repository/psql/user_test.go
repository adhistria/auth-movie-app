package psql_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/adhistria/auth-movie-app/infrastructure/storage"
	"github.com/adhistria/auth-movie-app/internal/domain"
	. "github.com/adhistria/auth-movie-app/internal/repository/psql"
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
	userRepo := NewUserRepository((&mockStorage))
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

func TestLoginUser(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error %s", err)
	}

	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mockUser := domain.User{
		ID:       1,
		Name:     "adhi satria",
		Email:    "adhistria1@gmail.com",
		Password: "password",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).AddRow(uint64(1), "adhi satria", "adhistria1@gmail.com", "password")

	mockDB.ExpectQuery(`SELECT email, password FROM users WHERE email = ?`).WithArgs(&mockUser.Email).WillReturnRows(rows)

	mockStorage := storage.Database{Conn: sqlxDB}
	userRepo := NewUserRepository((&mockStorage))
	user, err := userRepo.FindByEmail(context.Background(), &mockUser)
	if err != nil {
		t.Errorf("Error was not expected: %s", err)
	}
	t.Log(user)
}
