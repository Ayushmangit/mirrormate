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
		GetByID(context.Context, int64) (*User, error)
		GetByEmail(context.Context, string) (*User, error)
		UpdateByID(context.Context, int64, *User) (*User, error)
		DeleteByID(context.Context, int64) error
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
