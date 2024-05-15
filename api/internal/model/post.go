package model

import (
	"time"
)

type Post struct {
	ID              string    `db:"id" json:"id" validate:"required"`
	UserID          string    `db:"user_id" json:"user_id"`
	Title           string    `db:"title" json:"title" validate:"required"`
	Content         string    `db:"content" json:"content" validate:"required"`
	ImagePaths      []string  `db:"path" json:"image_path"`
	CommunityID     string    `db:"community_id" json:"community_id"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	Privacy         string    `json:"privacy" validate:"required|contains:private,public,selected"`
	UserInformation User      `json:"user_information"`
	GroupName       string    `json:"group_name"`
}

// NewPost creates a pointer to Post struct with new uuid
func NewPost() *Post {
	return &Post{
		CreatedAt: time.Now(),
	}
}
