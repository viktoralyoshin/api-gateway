package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viktoralyoshin/utils/pkg/errs"
)

func ReturnErr(c *fiber.Ctx, err error) error {
	code, msg := errs.HTTPStatus(err)

	return c.Status(code).JSON(fiber.Map{
		"error": msg,
	})
}
