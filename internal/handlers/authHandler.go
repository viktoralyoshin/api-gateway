package handlers

import (
	"api-gateway/internal/grpc"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	authpb "github.com/viktoralyoshin/playhub-proto/gen/go/auth"
)

func SignIn(c *fiber.Ctx) error {

	log.Info().Msg("/signin")

	return c.SendString("signin")
}

func SignUp(c *fiber.Ctx) error {
	ctx := c.UserContext()

	resp, err := grpc.AuthClient.Register(ctx, &authpb.RegisterRequest{Username: "username", Email: "fdsf", Password: "lsfd"})
	if err != nil {
		log.Error().Msgf("registration for user %s, %s failed: %v", "fd", "fds", err)

		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Info().Msgf("user %s has been registered: id=%s", "asda", resp.UserId)

	return c.Status(201).JSON(resp)
}
