package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/vitaly06/task-manager-rest-api/internal/service"
)

func AuthRequired(secret string) fiber.Handler {
	return func(c fiber.Ctx) error {
		tokenString := c.Cookies("access_token")

		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Вы не авторизованы",
			})
		}

		claims := &service.Claims{}

		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(t *jwt.Token) (any, error) {
				if t.Method != jwt.SigningMethodHS256 {
					return nil, jwt.ErrSignatureInvalid
				}

				return []byte(secret), nil
			},
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		)

		if err != nil || token == nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Недействительный или истёкший токен",
			})
		}

		if _, err := uuid.Parse(claims.UserID.String()); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Некорректный токен",
			})
		}

		c.Locals("userID", claims.UserID)

		return c.Next()
	}
}
