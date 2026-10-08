package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/vitaly06/task-manager-rest-api/internal/handler"
)

func AuthRoutes(router fiber.Router, handler *handler.AuthHandler) {
	router.Post("/auth/sign-up", handler.SignUp)
	router.Post("/auth/sign-in", handler.SignIn)
}
