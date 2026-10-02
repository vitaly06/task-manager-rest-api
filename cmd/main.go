package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/vitaly06/task-manager-rest-api/internal/config"
	"github.com/vitaly06/task-manager-rest-api/internal/consts"
	"github.com/vitaly06/task-manager-rest-api/internal/database"
)

func main() {
	cfg := config.NewConfig()

	_, err := database.ConnectDB(cfg.DSN)

	if err != nil {
		log.Fatalf("[%s] Не удалось подключиться к БД", consts.Red("ERROR"))
	}

	app := fiber.New()

	app.Listen(":" + cfg.Port)
}
