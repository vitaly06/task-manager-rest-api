package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/vitaly06/task-manager-rest-api/internal/config"
	"github.com/vitaly06/task-manager-rest-api/internal/consts"
	"github.com/vitaly06/task-manager-rest-api/internal/database"
	"github.com/vitaly06/task-manager-rest-api/internal/handler"
	"github.com/vitaly06/task-manager-rest-api/internal/repository"
	"github.com/vitaly06/task-manager-rest-api/internal/routes"
	"github.com/vitaly06/task-manager-rest-api/internal/service"
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

	api := app.Group("/api")

	// User
	userR := repository.NewUserRepository(db)

	// Auth
	authS := service.NewAuthService(userR, cfg.JwtSecret)
	authH := handler.NewAuthHandler(authS)

	routes.AuthRoutes(api, authH, cfg.JwtSecret)

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("[%s] Не удалось запустить сервер\n", consts.Red("ERROR"))
	}
}
