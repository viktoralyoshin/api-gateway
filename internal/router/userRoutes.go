package router

import (
	"api-gateway/internal/config"
	"api-gateway/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(r fiber.Router, cfg *config.Config) {
	auth := r.Group("/auth")

	authHandler := handlers.NewAuthHandler(cfg)

	auth.Post("/signin", authHandler.SignIn)
	auth.Post("/signup", authHandler.SignUp)
}
