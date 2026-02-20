package main

import (
	"context"
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users (name) VALUES ($1)RETURNING id`
	return r.db.QueryRowContext(ctx, query, user.Name).
		Scan(&user.ID)
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (*User, error) {
	query := `SELECT id, name FROM users WHERE id = $1`

	var user User

	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&user.ID, &user.Name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
