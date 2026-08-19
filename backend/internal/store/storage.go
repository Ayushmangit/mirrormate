package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrDuplicateEmail    = errors.New("New  payload")
	ErrDuplicateUsername = errors.New("record not found")
)

const QueryTimeoutDuration = time.Second * 5

type Storage struct {
	Users interface {
		Create(context.Context, *User) error
	}

	Roles interface {
		GetByName(context.Context, string) (*Role, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UserStorage{db},
		Roles: &RoleStorage{db},
	}
}
