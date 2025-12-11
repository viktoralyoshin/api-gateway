package middleware

import (
	"api-gateway/internal/grpc"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/viktoralyoshin/playhub-proto/gen/go/auth"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "authorization header required",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(401).JSON(fiber.Map{
				"error": "invalid token format",
			})
		}

		ctx := c.UserContext()

		tokenStr := parts[1]

		resp, err := grpc.AuthClient.ValidateToken(ctx, &auth.TokenRequest{TokenStr: tokenStr})
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		c.Locals("userId", resp.UserId)
		c.Locals("userRole", resp.UserRole)

		return c.Next()
	}
}
