package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func SignIn(c *fiber.Ctx) error {

	log.Info().Msg("/signin")

	return c.SendString("signin")
}

func SignUp(c *fiber.Ctx) error {

	log.Info().Msg("/signin")

	return c.SendString("signup")
}
