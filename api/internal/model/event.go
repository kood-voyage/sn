package model

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID          string    `db:"id" json:"id"`
	UserID      string    `db:"user_id" json:"user_id"`
	GroupID     string    `db:"group_id" json:"group_id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	Date        time.Time `db:"date" json:"date"`
}

func NewEvent() *Event {
	return &Event{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
	}
}
