package handler

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/vitaly06/task-manager-rest-api/internal/consts"
	"github.com/vitaly06/task-manager-rest-api/internal/service"
)

type AuthHandler struct {
	authS *service.AuthService
}

func NewAuthHandler(authS *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authS: authS,
	}
}

func (h *AuthHandler) SignUp(c fiber.Ctx) error {
	var req SignUpRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Почта и пароль обязательны для заполнения",
		})
	}

	err := h.authS.SignUp(req.Email, req.Password)

	if err != nil {
		if errors.Is(err, service.EmailAlreadyExists) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Пользователь с таким email уже существует",
			})
		}

		log.Printf("[%s] AuthService.SignUp: %s", consts.Red("ERROR"), err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Внутренняя ошибка сервера",
		})
	}

	log.Printf("[%s] Успешная регистрация: %s\n", consts.Green("INFO"), req.Email)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Регистрация прошла успешно",
	})
}

func (h *AuthHandler) SignIn(c fiber.Ctx) error {
	var req SignInRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	token, err := h.authS.SignIn(req.Email, req.Password)

	if err != nil {
		if errors.Is(err, service.PasswordsNotMatch) || errors.Is(err, service.EmailNotRegisted) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Неверный логин или пароль",
			})
		}

		log.Printf("[%s] AuthService.SignIn: %s", consts.Red("ERROR"), err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Внутренняя ошибка сервера",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		MaxAge:   24 * 60 * 60,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Успешная авторизация",
	})
}

func (h *AuthHandler) Logout(c fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
	})

	return c.SendStatus(fiber.StatusNoContent)
}
