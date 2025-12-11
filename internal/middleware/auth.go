package middleware

import (
	"api-gateway/internal/grpc"
	"api-gateway/internal/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/viktoralyoshin/playhub-proto/gen/go/auth"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Warn().Msg("request missing authorization header")

			return c.Status(401).JSON(fiber.Map{
				"error": "authorization header required",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Warn().Str("header", authHeader).Msg("invalid token format")

			return c.Status(401).JSON(fiber.Map{
				"error": "invalid token format",
			})
		}

		ctx := c.UserContext()
		tokenStr := parts[1]

		resp, err := grpc.AuthClient.ValidateToken(ctx, &auth.TokenRequest{TokenStr: tokenStr})
		if err != nil {
			log.Warn().Err(err).Msg("token validation failed via gRPC")

			return utils.ReturnErr(c, err)
		}

		c.Locals("userId", resp.UserId)
		c.Locals("userRole", resp.UserRole)

		return c.Next()
	}
}
