package router

import (
	"api-gateway/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(r fiber.Router) {
	auth := r.Group("/auth")

	auth.Get("/signin", handlers.SignIn)
	auth.Get("/signup", handlers.SignUp)
}
