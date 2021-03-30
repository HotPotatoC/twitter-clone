package action

import (
	"errors"
	"fmt"
	"strconv"

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
	followedID, err := strconv.ParseInt(c.Params("userID"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	if int64(followerID) == followedID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You cannot follow yourself",
		})
	}

	username, err := a.service.Execute(int64(followerID), followedID)
	if err != nil {
		if errors.Is(err, userEntity.ErrUserDoesNotExist) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}

		if errors.Is(err, entity.ErrUserAlreadyFollowed) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "You have already followed the user",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf("Followed %s", username),
	})
}
