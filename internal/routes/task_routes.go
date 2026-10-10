package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/vitaly06/task-manager-rest-api/internal/handler"
	"github.com/vitaly06/task-manager-rest-api/internal/middleware"
)

func TaskRoutes(router fiber.Router, handler *handler.TaskHandler, jwtSecret string) {
	router.Post("/tasks", middleware.AuthRequired(jwtSecret), handler.Create)
	router.Get("/tasks", middleware.AuthRequired(jwtSecret), handler.GetTasks)
	router.Get("/tasks/:id", middleware.AuthRequired(jwtSecret), handler.GetTaskByID)
	router.Put("/tasks/:id", middleware.AuthRequired(jwtSecret), handler.UpdateTask)
	router.Delete("/tasks/:id", middleware.AuthRequired(jwtSecret), handler.DeleteTask)
}
