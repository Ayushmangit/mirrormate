package store

import "database/sql"

type Role struct {
	ID    int64
	Name  string
	Level int
	Desc  string
}

type RoleStorage struct {
	db *sql.DB
}
