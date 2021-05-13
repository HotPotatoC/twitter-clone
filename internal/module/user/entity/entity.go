package entity

import "time"

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Handle    string    `json:"handle"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(id int64, name, handle, email, password string, createdAt time.Time) *User {
	return &User{
		ID:        id,
		Name:      name,
		Handle:    handle,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
	}
}
