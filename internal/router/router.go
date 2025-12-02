package router

import "github.com/gofiber/fiber/v2"

func SetupRouter(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	SetupUserRoutes(v1)
}
