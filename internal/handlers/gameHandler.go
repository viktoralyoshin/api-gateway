package handlers

import (
	"api-gateway/internal/grpc"
	"api-gateway/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	gamepb "github.com/viktoralyoshin/playhub-proto/gen/go/games"
)

type GameHandler struct{}

func NewGameHandler() *GameHandler {
	return &GameHandler{}
}

func (h *GameHandler) GetGame(c *fiber.Ctx) error {
	slug := c.Params("slug", "")
	if slug == "" {
		log.Fatal().Msg("GamesService GetGame: invalid param")

		return c.Status(500).JSON(fiber.Map{
			"error": "invalid param",
		})
	}

	ctx := c.UserContext()

	log.Info().Msgf("get game by slug: %s", slug)

	resp, err := grpc.GamesClient.GetGame(ctx, &gamepb.GetGameRequest{IdType: &gamepb.GetGameRequest_Slug{
		Slug: slug,
	}})
	if err != nil {
		log.Error().Msgf("get game %s failed: %v", slug, err)

		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	game := models.GameFromProto(resp.Game)

	return c.Status(200).JSON(game)
}

func (h *GameHandler) SearchGames(c *fiber.Ctx) error {
	query := c.Query("q", "")
	limitQuery := c.QueryInt("l", 10)

	ctx := c.UserContext()

	log.Info().Msgf("search game: %s", query)

	resp, err := grpc.GamesClient.SearchGames(ctx, &gamepb.SearchGamesRequest{Query: query, Limit: uint32(limitQuery)})
	if err != nil {
		log.Error().Msgf("search game %s failed: %v", query, err)

		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	gamesSlc := make([]models.Game, 0, len(resp.Games))

	for _, game := range resp.Games {
		gamesSlc = append(gamesSlc, *models.GameFromProto(game))
	}

	return c.Status(200).JSON(gamesSlc)
}

func (h *GameHandler) GetGamesByGenre(c *fiber.Ctx) error {
	return nil
}

func (h *GameHandler) GetTopRatedGames(c *fiber.Ctx) error {
	return nil
}

func (h *GameHandler) GetUpcomingGames(c *fiber.Ctx) error {
	return nil
}
