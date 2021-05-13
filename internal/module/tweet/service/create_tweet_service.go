package service

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/common/aws"
	"github.com/HotPotatoC/twitter-clone/internal/common/config"
	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/common/utils"
	"github.com/HotPotatoC/twitter-clone/internal/common/validator"
	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/pkg/errors"
)

type CreateTweetInput struct {
	Content string `json:"content" form:"content" validate:"required,max=300"`
}

func (i CreateTweetInput) Validate() []*validator.ValidationError {
	return validator.ValidateStruct(i)
}

type CreateTweetService interface {
	Execute(input CreateTweetInput, photos []*multipart.FileHeader, userID int64) error
}

type createTweetService struct {
	db database.Database
	s3 *aws.S3Bucket
}

func NewCreateTweetService(db database.Database, s3 *aws.S3Bucket) CreateTweetService {
	return createTweetService{db: db, s3: s3}
}

func (s createTweetService) Execute(input CreateTweetInput, photos []*multipart.FileHeader, userID int64) error {
	var photoURLs []string

	if len(photos) > 4 {
		return module.ErrTooManyAttachments
	}

	for _, photo := range photos {
		file, err := photo.Open()
		if err != nil {
			return errors.Wrap(err, "service.createTweetService.Execute")
		}
		defer file.Close()

		maxAttachmentSize := int64(config.GetInt("MAX_TWEET_ATTACHMENT_SIZE", 32*1024*1024))

		if photo.Size > maxAttachmentSize {
			return module.ErrUploadImageSizeTooLarge
		}

		buf := make([]byte, 512)
		_, err = file.Read(buf)
		if err != nil {
			return errors.Wrap(err, "service.createTweetService.Execute")
		}

		if !utils.IsValidImageContentType(http.DetectContentType(buf)) {
			return module.ErrInvalidImageType
		}

		fileKey := fmt.Sprintf("%d-%d-%s", userID, time.Now().Unix(), photo.Filename)

		out, err := s.s3.UploadObject(fileKey, file)
		if err != nil {
			return errors.Wrap(err, "service.createTweetService.Execute")
		}

		photoURLs = append(photoURLs, out.Location)
	}

	_, err := s.db.Exec("INSERT INTO tweets(content, id_user, photo_urls, created_at) VALUES($1, $2, $3, $4)",
		input.Content, userID, photoURLs, time.Now())
	if err != nil {
		return errors.Wrap(err, "service.createTweetService.Execute")
	}

	return nil
}
