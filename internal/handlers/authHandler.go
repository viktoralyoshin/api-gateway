package handlers

import (
	"api-gateway/internal/config"
	"api-gateway/internal/grpc"
	"api-gateway/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	authpb "github.com/viktoralyoshin/playhub-proto/gen/go/auth"
)

type AuthHandler struct {
	cfg *config.Config
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
	}
}

func (h *AuthHandler) SignIn(c *fiber.Ctx) error {

	ctx := c.UserContext()

	var body struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		log.Error().Msgf("failed to parse login body: %v", err)

		return c.Status(400).JSON(fiber.Map{
			"error": "invalid payload",
		})
	}

	log.Info().Msgf("login user: %s", body.Login)

	resp, err := grpc.AuthClient.Login(ctx, &authpb.LoginRequest{Login: body.Login, Password: body.Password})
	if err != nil {
		log.Error().Msgf("login for user %s failed: %v", body.Login, err)

		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Cookie(utils.SetRefreshToken(resp.RefreshToken, h.cfg.RefreshToketTtl, h.cfg.Domain))

	log.Info().Msgf("user %s has been logined: id=%s", body.Login, resp.UserId)

	return c.Status(200).JSON(fiber.Map{
		"id":           resp.UserId,
		"access_token": resp.AccessToken,
	})
}

func (h *AuthHandler) SignUp(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var body struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		log.Error().Msgf("failed to parse registration body: %v", err)

		return c.Status(400).JSON(fiber.Map{
			"error": "invalid payload",
		})
	}

	log.Info().Msgf("registration user: email=%s, username=%s", body.Email, body.Username)

	resp, err := grpc.AuthClient.Register(ctx, &authpb.RegisterRequest{Username: body.Username, Email: body.Email, Password: body.Password})
	if err != nil {
		log.Error().Msgf("registration for user %s, %s failed: %v", body.Email, body.Username, err)

		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Cookie(utils.SetRefreshToken(resp.RefreshToken, h.cfg.RefreshToketTtl, h.cfg.Domain))

	log.Info().Msgf("user %s has been registered: id=%s", body.Username, resp.UserId)

	return c.Status(201).JSON(fiber.Map{
		"id":           resp.UserId,
		"access_token": resp.AccessToken,
	})
}
