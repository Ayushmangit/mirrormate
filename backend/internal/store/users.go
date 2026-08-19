package store

import (
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type UserStorage struct {
	db *sql.DB
}

type User struct {
	ID        int64    `json:"id"`
	Email     string   `json:"email"`
	Password  password `json:"-"`
	Username  string   `json:"username"`
	CreatedAt string   `json:"created_at"`
	RoleID    int64    `json:"role_id"`
	IsActive  bool     `json:"is_active"`
	Role      Role     `json:"role"`
}

type password struct {
	text *string
	hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.text = &text
	p.hash = hash

	return nil
}

func (p *password) Compare(slug string) error {
	return bcrypt.CompareHashAndPassword(p.hash, []byte(slug))
}

func (s *UserStorage) Create(ctx context.Context, user *User) error {

	query := `
INSERT 
	INTO users(username,email,password,role_id) 
	values ($1,$2,$3,
	(SELECT id from roles where name = $4)) returning id,created_at;
`

	role := user.Role.Name
	if role == "" {
		role = "user"
	}

	err := s.db.QueryRowContext(ctx, query,
		user.Username,
		user.Email,
		user.Password.hash,
		role).Scan(
		&user.ID,
		&user.CreatedAt)
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

func (s *UserStorage) GetByID(ctx context.Context, userID int64) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	query := `SELECT id,username,email,created_at,role_id from users where id = $1`

	user := &User{}

	err := s.db.QueryRowContext(ctx, query, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.RoleID,
	)
	//TODO : error handling
	if err != nil {
		return nil, nil
	}
	return user, nil
}

func (s *UserStorage) GetByEmail(ctx context.Context, email string) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	query := `SELECT id,username,email,created_at,role_id from users where email = $1`

	user := &User{}

	err := s.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.RoleID,
	)
	//TODO: error handling
	if err != nil {
		return nil, nil
	}
	return user, nil
}

// TODO: complete update patch req
func (s *UserStorage) UpdateByID(ctx context.Context, userID int64, updatePayload *User) (*User, error) {

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	return nil, nil
}

func (s *UserStorage) DeleteByID(ctx context.Context, userID int64) error {

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `DELETE from users where id = $1`
	_, err := s.db.ExecContext(ctx, query, userID)
	//TODO: error handling
	if err != nil {
		return err
	}
	return nil
}
