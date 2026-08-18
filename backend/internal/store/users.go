package store

import (
	"context"
	"database/sql"
)

type UserStorage struct {
	db *sql.DB
}

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	RoleID    int64  `json:"role_id"`
	IsActive  bool   `json:"is_active"`
	Role      Role   `json:"role"`
}

func (s *UserStorage) Create(ctx context.Context, user *User) error {

	query := `
INSERT INTO users(username,email,password,role_id) values ($1,$2,$3,(SELECT id from roles where name = $4)) returning id,created_at;
`

	role := user.Role.Name
	if role == "" {
		role = "user"
	}

	err := s.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password, role).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "user_email_key"`:
			return ErrDuplicateEmail
		case err.Error() == `pq: duplicate key value violates unique constraint "user_username_key"`:
			return ErrDuplicateUsername
		default:
			return err
		}
	}

	return nil
}
