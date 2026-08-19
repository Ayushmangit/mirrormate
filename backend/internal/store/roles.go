package store

import (
	"context"
	"database/sql"
)

type Role struct {
	ID    int64
	Name  string
	Level int
	Desc  string
}

type RoleStorage struct {
	db *sql.DB
}

func (s *RoleStorage) GetByName(ctx context.Context, name string) (*Role, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `SELECT id,name,desc,level from roles where name = $1`
	role := &Role{}

	err := s.db.QueryRowContext(ctx,
		query,
		name).Scan(
		&role.ID,
		&role.Name,
		&role.Desc,
		&role.Level)
	if err != nil {
		return nil, err
	}
	return role, nil
}
