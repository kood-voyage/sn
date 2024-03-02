package model

import (
	"github.com/google/uuid"
)

type Group struct {
	ID          string `db:"id" json:"id" validate:"required"`
	CreatorID   string `db:"creator_id" json:"creator_id" validate:"required"`
	Name        string `db:"name" json:"name" validate:"required|min_len:2|max_len:25"`
	Description string `db:"description" json:"description" validate:"required"`
}

func NewGroup() *Group {
	id := uuid.New().String()
	return &Group{
		ID: id,
	}
}
