package repositories

import (
	"context"
	"offline-notes-app-backend/db"
)

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	return db.Pool.QueryRow(ctx, query, user.Username, user.Email, user.Password).Scan(&user.ID)
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*User, error) {
	user := &User{}
	query := `SELECT id, username, email, password FROM users WHERE id = $1`
	err := db.Pool.QueryRow(ctx, query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}
	query := `SELECT id, username, email, password FROM users WHERE email = $1`
	err := db.Pool.QueryRow(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *User) error {
	query := `UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4`
	_, err := db.Pool.Exec(ctx, query, user.Username, user.Email, user.Password, user.ID)
	return err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Pool.Exec(ctx, query, id)
	return err
}

func (r *UserRepository) AuthenticateUser(ctx context.Context, email, password string) (*User, error) {
	user := &User{}
	query := `SELECT id, username, email, password FROM users WHERE email = $1 AND password = $2`
	err := db.Pool.QueryRow(ctx, query, email, password).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
