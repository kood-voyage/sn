package model

import "github.com/google/uuid"

type Follower struct {
	ID       string `db:"id" json:"id" validate:"required"`
	SourceID string `db:"source_id" json:"source_id" validate:"required"`
	TargetID string `db:"target_id" json:"target_id" validate:"required"`
}

func NewFollow() *Follower {
	id := uuid.New().String()
	return &Follower{
		ID: id,
	}
}
