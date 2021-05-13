package action

import (
	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/service"
	"github.com/gofiber/fiber/v2"
)

type createTweetAction struct {
	service service.CreateTweetService
}

func NewCreateTweetAction(service service.CreateTweetService) module.Action {
	return createTweetAction{service: service}
}

func (a createTweetAction) Execute(c *fiber.Ctx) error {
	var input service.CreateTweetInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errors := input.Validate(); errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	mf, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userID := c.Locals("userID").(float64)

	err = a.service.Execute(input, mf.File["photos"], int64(userID))
	if err != nil {
		switch err {
		case module.ErrTooManyAttachments:
			return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{
				"message": "Only a maximum of 4 images are allowed",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "There was a problem on our side",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully posted a new tweet",
	})
}
