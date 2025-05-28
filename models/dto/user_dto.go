package dto

// RegisterUserRequest is the payload for registering a new user
// Constraints: name, email, and password are required
// id is not provided by the client
// email should be validated as an email in handler/service

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// RegisterUserResponse is the response after successful registration
// id, name, email are returned; password is omitted for security
// created_at, updated_at, deleted_at are optional (nullable)
type RegisterUserResponse struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}
