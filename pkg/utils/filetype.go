package utils

var ImageTypes = []string{"image/jpeg", "image/jpg", "image/png"}

func IsValidImageContentType(contentType string) bool {
	for _, it := range ImageTypes {
		if it == contentType {
			return true
		}
	}

	return false
}
