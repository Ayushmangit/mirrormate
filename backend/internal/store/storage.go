package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrDuplicateEmail    = errors.New("record not found")
	ErrDuplicateUsername = errors.New("record not found")
	ErrNotFound          = errors.New("resource not found")
	ErrUnAuthorized      = errors.New("unauthorized")
	ErrNotActivated      = errors.New("user not activated")
	ErrBadRequest        = errors.New("bad request")
)

const QueryTimeoutDuration = time.Second * 5

type Storage struct {
	Users interface {
		Create(context.Context, *User) error
		CreateAndInvite(context.Context, *User, string, time.Duration) error
		Activate(context.Context, string) error
		GetByID(context.Context, int64) (*User, error)
		GetByEmail(context.Context, string) (*User, error)
		UpdateByID(context.Context, int64, UpdateUserPayload) (*User, error)
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

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
