package services

import (
	"context"
	"errors"
	"offline-notes-app-backend/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// Register a new user (add password hashing in real app)
func (s *UserService) RegisterUser(ctx context.Context, user *repositories.User) error {
	// TODO: Add validation, check for existing email, hash password
	return s.Repo.CreateUser(ctx, user)
}

// Authenticate user by email and password
func (s *UserService) AuthenticateUser(ctx context.Context, email, password string) (*repositories.User, error) {
	user, err := s.Repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

// Get user by ID
func (s *UserService) GetUserByID(ctx context.Context, id int64) (*repositories.User, error) {
	return s.Repo.GetUserByID(ctx, id)
}

// Update user
func (s *UserService) UpdateUser(ctx context.Context, user *repositories.User) error {
	return s.Repo.UpdateUser(ctx, user)
}

// Delete user
func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.Repo.DeleteUser(ctx, id)
}
