package module

import "errors"

var (
	ErrInvalidCursor = errors.New("invalid cursor")

	ErrUploadImageSizeTooLarge = errors.New("upload image size too large")
	ErrInvalidImageType        = errors.New("invalid image type")
	ErrTooManyAttachments      = errors.New("too many attachments")
)
