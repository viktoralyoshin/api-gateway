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
		log.Error().Err(err).Msg("failed to parse login body")

		return c.Status(400).JSON(fiber.Map{
			"error": "invalid payload",
		})
	}

	log.Info().Str("login", body.Login).Msg("login attempt")

	resp, err := grpc.AuthClient.Login(ctx, &authpb.LoginRequest{Login: body.Login, Password: body.Password})
	if err != nil {
		log.Error().
			Err(err).
			Str("login", body.Login).
			Msg("login failed via gRPC")

		return utils.ReturnErr(c, err)
	}

	c.Cookie(utils.SetRefreshToken(resp.RefreshToken, h.cfg.RefreshToketTtl, h.cfg.Domain))

	log.Info().
		Str("login", body.Login).
		Str("user_id", resp.UserId).
		Msg("user successfully logged in")

	return c.Status(200).JSON(fiber.Map{
		"id":           resp.UserId,
		"username":     resp.Username,
		"email":        resp.Email,
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
		log.Error().Err(err).Msg("failed to parse registration body")

		return c.Status(400).JSON(fiber.Map{
			"error": "invalid payload",
		})
	}

	log.Info().
		Str("email", body.Email).
		Str("username", body.Username).
		Msg("registration attempt")

	resp, err := grpc.AuthClient.Register(ctx, &authpb.RegisterRequest{Username: body.Username, Email: body.Email, Password: body.Password})
	if err != nil {
		log.Error().
			Err(err).
			Str("email", body.Email).
			Str("username", body.Username).
			Msg("registration failed via gRPC")

		return utils.ReturnErr(c, err)
	}

	c.Cookie(utils.SetRefreshToken(resp.RefreshToken, h.cfg.RefreshToketTtl, h.cfg.Domain))

	log.Info().
		Str("username", body.Username).
		Str("user_id", resp.UserId).
		Msg("user successfully registered")

	return c.Status(201).JSON(fiber.Map{
		"id":           resp.UserId,
		"username":     resp.Username,
		"email":        resp.Email,
		"access_token": resp.AccessToken,
	})
}
