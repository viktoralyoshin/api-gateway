package app

import (
	"api-gateway/internal/config"
	"api-gateway/internal/grpc"
	"api-gateway/internal/router"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Start(cfg *config.Config) {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &log.Logger,
	}))

	grpc.Init(cfg)

	router.SetupRouter(app, cfg)

	log.Info().Msg("Api Gateway starting :8080")

	if err := app.Listen(":8080"); err != nil {
		log.Fatal().Err(err).Msg("Api Gateway starting failed")
	}
}
