package action

import (
	"fmt"

	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/service"
	"github.com/gofiber/fiber/v2"
)

type searchTweetAction struct {
	service service.SearchTweetService
}

func NewSearchTweetAction(service service.SearchTweetService) module.Action {
	return searchTweetAction{service: service}
}

func (a searchTweetAction) Execute(c *fiber.Ctx) error {
	searchQuery := c.Query("query")
	if searchQuery == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Missing search query",
		})
	}

	userID := c.Locals("userID").(float64)
	cursor := c.Query("cursor")
	tweets, err := a.service.Execute(searchQuery, int64(userID), cursor)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total_items": len(tweets),
		"items":       tweets,
	})
}
