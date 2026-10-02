package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/vitaly06/task-manager-rest-api/internal/config"
)

func main() {
	cfg := config.NewConfig()

	app := fiber.New()

	app.Listen(":" + cfg.Port)
}
