package repository

import (
	"context"
	"database/sql"
	"go-gin-auth/domain"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{db}
}

func (r *Users) Create(ctx context.Context, user domain.User) error {
	return nil
}

func (r *Users) GetByCredentials(ctx context.Context, email, password string) (domain.User, error) {
	var user domain.User
	selectStmt := "SELECT id, name, surname, email, registered_at FROM users WHERE email=$1 AND password=$2"
	err := r.db.QueryRowContext(ctx, selectStmt, email, password).
		Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.RegisteredAt)

	return user, err
}
