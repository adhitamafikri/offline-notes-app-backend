package routes

import (
	"github.com/gofiber/fiber/v2"
	"offline-notes-app-backend/handlers"
)

func RegisterRoutes(app *fiber.App) {
	// Notes routes
	app.Get("/api/notes", handlers.getNotes)          // List all notes
	app.Post("/api/notes", handlers.createNote)       // Create a new note
	app.Get("/api/notes/:id", handlers.getNote)       // Get a note by ID
	app.Put("/api/notes/:id", handlers.updateNote)    // Update a note by ID
	app.Delete("/api/notes/:id", handlers.deleteNote) // Delete a note by ID

	// UserSettings routes
	app.Get("/api/user-settings", handlers.getUserSettings)    // Get user settings
	app.Put("/api/user-settings", handlers.updateUserSettings) // Update user settings

	// User routes
	app.Post("/api/users/register", handlers.registerUser) // Register a new user
	app.Post("/api/users/login", handlers.loginUser)       // User login
	app.Get("/api/users/:id", handlers.getUser)            // Get user by ID
	app.Put("/api/users/:id", handlers.updateUser)         // Update user by ID
	app.Delete("/api/users/:id", handlers.deleteUser)      // Delete user by ID
}
