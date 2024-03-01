package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserWithChat struct {
	User   User   `json:"user"`
	ChatID string `json:"chat_id"`
}

type User struct {
	ID          string    `db:"id" json:"id"`
	Username    string    `db:"username" json:"username" validate:"min_len:4|max_len:32"`
	Email       string    `db:"email" json:"email" validate:"required|email"`
	Password    string    `db:"password" json:"password" validate:"min_len:8|max_len:32"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	DateOfBirth string    `db:"date_of_birth" json:"date_of_birth"`
	FirstName   string    `db:"first_name" json:"first_name" validate:"min_len:2|max_len:32"`
	LastName    string    `db:"last_name" json:"last_name" validate:"min_len:2|max_len:32"`
	Gender      string    `db:"gender" json:"gender" validate:"required"`
	Description string    `db:"description" json:"description"`
	ImageURL    string    `db:"image_url" json:"image_url"`
}

// Sanitize erase password so it would not appear in respond
func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func encryptString(str string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.Password = enc
	}

	return nil
}
