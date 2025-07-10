package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ApiPayments(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "OK"})

}
