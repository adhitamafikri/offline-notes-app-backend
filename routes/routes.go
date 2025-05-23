package routes

import (
	"github.com/gofiber/fiber/v2"
	"offline-notes-app-backend/handlers"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Notes routes
	v1.Get("/notes", handlers.GetNotes)          // List all notes
	v1.Post("/notes", handlers.CreateNote)       // Create a new note
	v1.Get("/notes/:id", handlers.GetNote)       // Get a note by ID
	v1.Put("/notes/:id", handlers.UpdateNote)    // Update a note by ID
	v1.Delete("/notes/:id", handlers.DeleteNote) // Delete a note by ID

	// UserSettings routes
	v1.Get("/user-settings", handlers.GetUserSettings)    // Get user settings
	v1.Put("/user-settings", handlers.UpdateUserSettings) // Update user settings

	// User routes
	v1.Post("/users/register", handlers.RegisterUser) // Register a new user
	v1.Post("/users/login", handlers.LoginUser)       // User login
	v1.Get("/users/:id", handlers.GetUser)            // Get user by ID
	v1.Put("/users/:id", handlers.UpdateUser)         // Update user by ID
	v1.Delete("/users/:id", handlers.DeleteUser)      // Delete user by ID
}
