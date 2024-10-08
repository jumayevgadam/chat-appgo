package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContent(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertedId int
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&lastInsertedId)
	if err != nil {
		return &User{}, nil
	}

	user.ID = int64(lastInsertedId)
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	query := "SELECT id, email, username, password FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	if err != nil {
		return &User{}, nil
	}

	return &u, nil
}
