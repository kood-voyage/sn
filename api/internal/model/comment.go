package model

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID              string    `db:"id" json:"id" validate:"required"`
	UserID          string    `db:"user_id" json:"user_id" `
	PostID          string    `db:"post_id" json:"post_id" validate:"required"`
	ParentID        string    `db:"parent_id" json:"parent_id"`
	Content         string    `db:"content" json:"content" validate:"required"`
	ImagePaths      []string  `db:"path" json:"image_path"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UserName        string    `db:"user_name" json:"user_name"`
	UserAvatar      string    `db:"user_avatar" json:"user_avatar"`
	Count           string    `json:"count"`
	UserInformation User      `json:"user_information"`
}

// NewComment creates a pointer to Comment struct with new uuid
func NewComment() *Comment {
	id := uuid.New().String()
	return &Comment{
		ID:        id,
		CreatedAt: time.Now(),
	}
}
