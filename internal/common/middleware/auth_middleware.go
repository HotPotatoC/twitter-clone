package middleware

import (
	"strings"

	"github.com/HotPotatoC/twitter-clone/internal/common/token"
	"github.com/gofiber/fiber/v2"
)

type authMiddleware struct{}

func NewAuthMiddleware() Middleware {
	return authMiddleware{}
}

func (m authMiddleware) Execute() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var accessToken string
		if c.Cookies("access_token") != "" {
			accessToken = c.Cookies("access_token")
		} else {
			authorization := c.Get("Authorization")
			if len(strings.Split(authorization, " ")) < 2 {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Please login to continue",
				})
			}
			accessToken = strings.Split(authorization, " ")[1]
		}

		claims, err := token.VerifyAccessToken(accessToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Please login to continue",
			})
		}

		userID, ok := claims["userID"]
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Please try and login again",
			})
		}

		handle, ok := claims["handle"]
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Please try and login again",
			})
		}

		email, ok := claims["email"]
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Please try and login again",
			})
		}

		c.Locals("userID", userID)
		c.Locals("handle", handle)
		c.Locals("email", email)
		c.Locals("accessToken", accessToken)

		return c.Next()
	}
}
