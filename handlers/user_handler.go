package handlers

import (
	"context"
	"offline-notes-app-backend/models/dto"
	"offline-notes-app-backend/repositories"
	"offline-notes-app-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var userService *services.UserService

func InitUserHandler(repo *repositories.UserRepository) {
	userService = services.NewUserService(repo)
}

func RegisterUser(c *fiber.Ctx) error {
	var req dto.RegisterUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	// Map DTO to repository model
	user := &repositories.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := userService.RegisterUser(context.Background(), user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// Prepare response DTO
	resp := dto.RegisterUserResponse{
		ID:        strconv.FormatInt(user.ID, 10),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: nil, // Set if available
		UpdatedAt: nil, // Set if available
		DeletedAt: nil, // Set if available
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func LoginUser(c *fiber.Ctx) error {
	input := new(struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	})
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	user, err := userService.AuthenticateUser(context.Background(), input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	user, err := userService.GetUserByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	user := new(repositories.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	user.ID = id
	if err := userService.UpdateUser(context.Background(), user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	if err := userService.DeleteUser(context.Background(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
