package app

import (
	"api-gateway/internal/router"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Start() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &log.Logger,
	}))

	router.SetupRouter(app)

	log.Info().Msg("Api Gateway starting :8000")

	if err := app.Listen(":8000"); err != nil {
		log.Fatal().Err(err).Msg("Api Gateway starting failed")
	}
}
