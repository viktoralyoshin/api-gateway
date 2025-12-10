package router

import (
	"api-gateway/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupGamesRoutes(r fiber.Router) {
	games := r.Group("/games")

	gameHandler := handlers.NewGameHandler()

	games.Get("/search", gameHandler.SearchGames)
	games.Get("/top", gameHandler.GetTopRatedGames)
	games.Get("/genre/:slug", gameHandler.GetGamesByGenre)
	games.Get("/:slug", gameHandler.GetGame)
	games.Get("/upcoming", gameHandler.GetUpcomingGames)
}
