package module

import "github.com/gofiber/fiber/v2"

type Action interface {
	Execute(c *fiber.Ctx) error
}
