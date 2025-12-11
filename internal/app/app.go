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

	// Подключаем логгер для Fiber (он будет логировать каждый запрос)
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &log.Logger,
	}))

	grpc.Init(cfg)

	router.SetupRouter(app, cfg)

	log.Info().
		Str("port", "8080").
		Msg("API Gateway starting")

	if err := app.Listen(":8080"); err != nil {
		log.Fatal().Err(err).Msg("API Gateway server failed")
	}
}
