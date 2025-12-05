package router

import (
	"api-gateway/internal/config"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, cfg *config.Config) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	SetupUserRoutes(v1, cfg)
}
