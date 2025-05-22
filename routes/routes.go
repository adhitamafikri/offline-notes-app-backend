package routes

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	// Notes routes
	app.Get("/api/notes", getNotes)          // List all notes
	app.Post("/api/notes", createNote)       // Create a new note
	app.Get("/api/notes/:id", getNote)       // Get a note by ID
	app.Put("/api/notes/:id", updateNote)    // Update a note by ID
	app.Delete("/api/notes/:id", deleteNote) // Delete a note by ID

	// UserSettings routes
	app.Get("/api/user-settings", getUserSettings)    // Get user settings
	app.Put("/api/user-settings", updateUserSettings) // Update user settings

	// User routes
	app.Post("/api/users/register", registerUser) // Register a new user
	app.Post("/api/users/login", loginUser)       // User login
	app.Get("/api/users/:id", getUser)            // Get user by ID
	app.Put("/api/users/:id", updateUser)         // Update user by ID
	app.Delete("/api/users/:id", deleteUser)      // Delete user by ID
}

// Handler function stubs (implement these in your handlers package or here for now)
func getNotes(c *fiber.Ctx) error           { return c.SendStatus(fiber.StatusNotImplemented) }
func createNote(c *fiber.Ctx) error         { return c.SendStatus(fiber.StatusNotImplemented) }
func getNote(c *fiber.Ctx) error            { return c.SendStatus(fiber.StatusNotImplemented) }
func updateNote(c *fiber.Ctx) error         { return c.SendStatus(fiber.StatusNotImplemented) }
func deleteNote(c *fiber.Ctx) error         { return c.SendStatus(fiber.StatusNotImplemented) }
func getUserSettings(c *fiber.Ctx) error    { return c.SendStatus(fiber.StatusNotImplemented) }
func updateUserSettings(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusNotImplemented) }
func registerUser(c *fiber.Ctx) error       { return c.SendStatus(fiber.StatusNotImplemented) }
func loginUser(c *fiber.Ctx) error          { return c.SendStatus(fiber.StatusNotImplemented) }
func getUser(c *fiber.Ctx) error            { return c.SendStatus(fiber.StatusNotImplemented) }
func updateUser(c *fiber.Ctx) error         { return c.SendStatus(fiber.StatusNotImplemented) }
func deleteUser(c *fiber.Ctx) error         { return c.SendStatus(fiber.StatusNotImplemented) }
