package model

import (
	"time"

	"github.com/google/uuid"
)

type Request struct {
	ID        string    `db:"id" json:"id" validate:"required"`
	TypeID    int       `db:"type_id" json:"type_id" validate:"required"`
	SourceID  string    `db:"source_id" json:"source_id" validate:"required"`
	TargetID  string    `db:"target_id" json:"target_id" validate:"required"`
	Message   string    `db:"message" json:"message"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// type RequestType struct {
// 	ID          string `db:"id" json:"id"`
// 	Description string `db:"description" json:"description"`
// }

func NewRequest() *Request {
	id := uuid.New().String()
	return &Request{
		ID: id,
	}
}
