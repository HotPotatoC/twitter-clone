package auth

import "github.com/gofiber/fiber/v2"

func Routes(r *fiber.App) {
	r.Post("/login")
	r.Post("/logout")
	r.Get("/me")
}
