package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/vitaly06/task-manager-rest-api/internal/handler"
	"github.com/vitaly06/task-manager-rest-api/internal/middleware"
)

func AuthRoutes(router fiber.Router, handler *handler.AuthHandler, jwtSecret string) {
	router.Post("/auth/sign-up", handler.SignUp)
	router.Post("/auth/sign-in", handler.SignIn)
	router.Post("/auth/logout", middleware.AuthRequired(jwtSecret), handler.Logout)
}
