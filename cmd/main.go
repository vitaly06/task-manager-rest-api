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

	db, err := database.ConnectDB(cfg.DSN)

	if err != nil {
		log.Fatalf("[%s] Не удалось подключиться к БД\n", consts.Red("ERROR"))
	}

	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("[%s] Не удалось выполнить автомиграцию\n", consts.Red("ERROR"))
	}

	app := fiber.New()

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("[%s] Не удалось запустить сервер\n", consts.Red("ERROR"))
	}
}
