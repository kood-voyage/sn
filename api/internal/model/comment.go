package model

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ID        string    `db:"id" json:"id" validate:"required"`
	UserID    string    `db:"user_id" json:"user_id" validate:"required"`
	PostID    string    `db:"post_id" json:"post_id" validate:"required"`
	ParentID  string    `db:"parent_id" json:"parent_id"`
	Content   string    `db:"content" json:"content" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// NewComment creates a pointer to Comment struct with new uuid
func NewComment() *Comment {
	id := uuid.New().String()
	return &Comment{
		ID: id,
	}
}
