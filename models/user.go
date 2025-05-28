package models

// User represents a user in the system.
type User struct {
	ID        string  `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	Email     string  `json:"email" db:"email"`
	Password  string  `json:"password" db:"password"`
	CreatedAt *string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *string `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty" db:"deleted_at"`
}
