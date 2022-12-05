package models

import (
	"errors"
	"net/mail"
	"time"
	"unicode"

	userpb "github.com/HotPotatoC/twitter-clone/user/rpc/user"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// User represents a generic user
type User struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	ScreenName       string    `json:"screen_name"`
	PasswordHash     string    `json:"password_hash"`
	Email            string    `json:"email"`
	Bio              string    `json:"bio"`
	Location         string    `json:"location"`
	Website          string    `json:"website"`
	ProfileImageURL  string    `json:"profile_image_url"`
	ProfileBannerURL string    `json:"profile_banner_url"`
	BirthDate        time.Time `json:"birth_date"`
	FollowersCount   int       `json:"followers_count"`
	FollowingsCount  int       `json:"followings_count"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (u User) PB() *userpb.User {
	return &userpb.User{
		UserId:           u.ID,
		Name:             u.Name,
		ScreenName:       u.ScreenName,
		Email:            u.Email,
		Bio:              u.Bio,
		Location:         u.Location,
		Website:          u.Website,
		BirthDate:        timestamppb.New(u.BirthDate),
		ProfileImageUrl:  u.ProfileImageURL,
		ProfileBannerUrl: u.ProfileBannerURL,
		FollowersCount:   int32(u.FollowersCount),
		FollowingsCount:  int32(u.FollowingsCount),
		CreatedAt:        timestamppb.New(u.CreatedAt),
		UpdatedAt:        timestamppb.New(u.UpdatedAt),
	}
}

// Validate validates the user fields
func (u User) Validate() error {
	if u.ScreenName == "" {
		return errors.New("missing User.ScreenName")
	}

	if u.Email == "" {
		return errors.New("missing User.Email")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("invalid email address")
	}

	return nil
}

// Password represents a password
type Password string

// Validate validates the password
func (p Password) Validate() error {
	if !p.validateCharacters() {
		return errors.New("invalid password characters")
	}

	return nil
}

// validateCharacters validates the password characters
func (p Password) validateCharacters() bool {
	for _, char := range p {
		if char > unicode.MaxASCII {
			return false
		}
	}

	return true
}

// GenerateHash generates a hash for the password
func (p Password) GenerateHash() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Compare compares the password with the hash
func (p Password) Compare(hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	if err != nil {
		return false, err
	}

	return true, nil
}

// PasswordIsValid determines wether the given password matches with the user's password hash
func (u User) PasswordIsValid(password Password) (bool, error) {
	return password.Compare(u.PasswordHash)
}
