package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          string    `db:"id" json:"id" validate:"required"`
	Username    string    `db:"username" json:"username" validate:"required"`
	Email       string    `db:"email" json:"email" validate:"required|email"`
	Password    string    `db:"password" json:"password" validate:"required|min_len:8"`
	CreatedAt   time.Time `db:"timestamp" json:"timestamp"`
	DateOfBirth string    `db:"date_of_birth" json:"date_of_birth"`
	FirstName   string    `db:"first_name" json:"first_name"`
	LastName    string    `db:"last_name" json:"last_name"`
	Description string    `db:"description" json:"description"`
	Avatar      string    `db:"avatar" json:"avatar"`
	Cover       string    `db:"cover" json:"cover"`
	Privacy     string    `json:"privacy" validate:"required|contains:public,private"`
	MemberType  int       `json:"member_type"`
	EventStatus string    `json:"event_status"`
}

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
