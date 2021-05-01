package service

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/aws"
	"github.com/HotPotatoC/twitter-clone/internal/config"
	"github.com/HotPotatoC/twitter-clone/internal/database"
	"github.com/HotPotatoC/twitter-clone/internal/utils"
	"github.com/pkg/errors"
)

type UpdateProfileImageService interface {
	Execute(photo *multipart.FileHeader, userID int64) (string, error)
}

type updateProfileImageService struct {
	db database.Database
	s3 *aws.S3Bucket
}

func NewUpdateProfileImageService(db database.Database, s3 *aws.S3Bucket) UpdateProfileImageService {
	return updateProfileImageService{db: db, s3: s3}
}

func (s updateProfileImageService) Execute(photo *multipart.FileHeader, userID int64) (string, error) {
	photoFile, err := photo.Open()
	if err != nil {
		return "", errors.Wrap(err, "service.updateProfileImageService.Execute")
	}
	defer photoFile.Close()

	maxUploadSize := int64(config.GetInt("MAX_UPLOAD_SIZE", 2.5*1024*1024))

	if photo.Size > maxUploadSize {
		return "", ErrUploadImageSizeTooLarge
	}

	buf := make([]byte, 512)
	_, err = photoFile.Read(buf)
	if err != nil {
		return "", errors.Wrap(err, "service.updateProfileImageService.Execute")
	}

	if !utils.IsValidImageContentType(http.DetectContentType(buf)) {
		return "", ErrInvalidImageType
	}

	fileKey := fmt.Sprintf("%d-%d-%s", userID, time.Now().Unix(), photo.Filename)

	out, err := s.s3.UploadObject(fileKey, photoFile)
	if err != nil {
		return "", errors.Wrap(err, "service.updateProfileImageService.Execute")
	}

	photoURL := out.Location

	_, err = s.db.Exec("UPDATE users SET photo_url = $1 WHERE id = $2", photoURL, userID)
	if err != nil {
		return "", errors.Wrap(err, "service.updateProfileImageService.Execute")
	}

	return photoURL, nil
}
