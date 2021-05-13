package action

import (
	"errors"
	"fmt"

	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/relationship/entity"
	"github.com/HotPotatoC/twitter-clone/internal/module/relationship/service"
	userEntity "github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/gofiber/fiber/v2"
)

type followUserAction struct {
	service service.FollowUserService
}

func NewFollowUserAction(service service.FollowUserService) module.Action {
	return followUserAction{service: service}
}

func (a followUserAction) Execute(c *fiber.Ctx) error {
	followerID := c.Locals("userID").(float64)
	followedID, err := c.ParamsInt("userID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
		})
	}

	if int(followerID) == followedID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You cannot follow yourself",
		})
	}

	username, err := a.service.Execute(int64(followerID), int64(followedID))
	if err != nil {
		switch {
		case errors.Is(err, userEntity.ErrUserDoesNotExist):
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		case errors.Is(err, entity.ErrUserAlreadyFollowed):
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "You have already followed the user",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "There was a problem on our side",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf("Followed %s", username),
	})
}
