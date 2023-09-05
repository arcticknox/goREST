package handlers

import (
	"goREST/models"
	"goREST/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	var userInput *models.UserRequest = new(models.UserRequest)
	// Parse req body
	err := c.BodyParser(userInput)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	token, err := services.SignUp(*userInput)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response[any]{
		Success: true,
		Message: "token",
		Data:    token,
	})
}

func Login(c *fiber.Ctx) error {
	var userInput *models.UserRequest = new(models.UserRequest)
	// Parse req body
	err := c.BodyParser(userInput)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	token, err := services.Login(*userInput)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response[any]{
		Success: true,
		Message: "token",
		Data:    token,
	})

}
