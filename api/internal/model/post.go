package model

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID          string    `db:"id" json:"id" validate:"required"`
	UserID      string    `db:"user_id" json:"user_id"`
	Title       string    `db:"title" json:"title" validate:"required"`
	Content     string    `db:"content" json:"content" validate:"required"`
	CommunityID string    `db:"community_id" json:"community_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	Privacy     string    `json:"privacy" validate:"required|contains:private,public,selected"`
}

// NewPost creates a pointer to Post struct with new uuid
func NewPost() *Post {
	id := uuid.New().String()
	return &Post{
		ID:        id,
		CreatedAt: time.Now(),
	}
}
