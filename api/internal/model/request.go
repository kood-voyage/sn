package model

import (
	"time"

	"github.com/google/uuid"
)

type Request struct {
	ID                string    `db:"id" json:"id" validate:"required"`
	TypeID            int       `db:"type_id" json:"type_id" validate:"required"`
	SourceID          string    `db:"source_id" json:"source_id"`
	TargetID          string    `db:"target_id" json:"target_id" validate:"required"`
	ParentID          string    `db:"parent_id" json:"parent_id"`
	Message           string    `db:"message" json:"message"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	SourceInformation User      `json:"source_information"`
	TargetInformation User      `json:"target_information"`
}

func FollowRequest() *Request {
	id := uuid.New().String()
	return &Request{
		ID:     id,
		TypeID: InitializeTypes().Request.Follow,
	}
}

func NotificationRequest() *Request {
	id := uuid.New().String()
	return &Request{
		ID:     id,
		TypeID: InitializeTypes().Request.Notification,
	}
}

func InviteRequest() *Request {
	id := uuid.New().String()
	return &Request{
		ID:     id,
		TypeID: InitializeTypes().Request.Invite,
	}
}
