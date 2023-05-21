package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"generate-id-use-autoincrement/internal/domain/user"
)

type userRepositoryInSqlite struct {
	db *sql.DB
}

func NewUserRepositoryInSqlite(db *sql.DB) user.Repository {
	return &userRepositoryInSqlite{db: db}
}

func (r *userRepositoryInSqlite) FindById(_ context.Context, userId *user.Id) (*user.User, error) {

	stmt, err := r.db.Prepare("SELECT name FROM users WHERE id = ?")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fail to find user: %v", err))
	}
	defer stmt.Close()

	var name string
	err = stmt.QueryRow(userId.Value()).Scan(&name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fail to find user: %v", err))
	}

	foundUser, err := user.ReconstructUser(userId.Value(), name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fail to find user: %v", err))
	}

	return foundUser, nil
}

func (r *userRepositoryInSqlite) Save(_ context.Context, user *user.User) error {

	tx, err := r.db.Begin()
	if err != nil {
		return errors.New(fmt.Sprintf("fail to save user: %v", err))
	}

	stmt, err := tx.Prepare("INSERT INTO users(name) VALUES (?)")
	if err != nil {
		_ = tx.Rollback()
		return errors.New(fmt.Sprintf("fail to save user: %v", err))
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name())
	if err != nil {
		_ = tx.Rollback()
		return errors.New(fmt.Sprintf("fail to save user: %v", err))
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return errors.New(fmt.Sprintf("fail to save user: %v", err))
	}

	return nil
}
