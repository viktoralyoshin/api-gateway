package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetRefreshToken(token string, expires time.Duration, domain string) *fiber.Cookie {
	cookie := new(fiber.Cookie)

	cookie.Name = "refresh_token"
	cookie.Value = token
	cookie.HTTPOnly = true
	cookie.Domain = domain
	cookie.Path = "/"
	cookie.SameSite = "none"
	cookie.Secure = true
	cookie.Expires = time.Now().Add(expires)

	return cookie
}
