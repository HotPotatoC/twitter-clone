package action

import (
	"errors"
	"fmt"

	"github.com/HotPotatoC/twitter-clone/internal/common/config"
	"github.com/HotPotatoC/twitter-clone/internal/common/utils"
	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/service"
	"github.com/gofiber/fiber/v2"
)

type updateProfileImageAction struct {
	service service.UpdateProfileImageService
}

func NewUpdateProfileImageAction(service service.UpdateProfileImageService) module.Action {
	return updateProfileImageAction{service: service}
}

func (a updateProfileImageAction) Execute(c *fiber.Ctx) error {
	photo, err := c.FormFile("photo")
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Missing image",
		})
	}

	userID := c.Locals("userID").(float64)
	photoURL, err := a.service.Execute(photo, int64(userID))
	if err != nil {
		switch {
		case errors.Is(err, module.ErrInvalidImageType):
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": "Invalid image type",
				"allowed": utils.ImageTypes,
			})
		case errors.Is(err, module.ErrUploadImageSizeTooLarge):
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": fmt.Sprintf("Image size is too big [Max: %s]",
					utils.ByteCount(int64(config.GetInt("MAX_UPLOAD_SIZE", 2.5*1024*1024)))),
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "There was a problem on our side",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"messages":  "Updated user profile image",
		"photo_url": photoURL,
	})
}
