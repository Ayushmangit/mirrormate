package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrDuplicateEmail    = errors.New("New  payload")
	ErrDuplicateUsername = errors.New("record not found")
)

type Storage struct {
	Users interface {
		Create(context.Context, *User) error
	}

	Roles interface {
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UserStorage{db},
		Roles: &RoleStorage{db},
	}
}
