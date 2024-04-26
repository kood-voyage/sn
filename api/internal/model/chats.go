package model

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID      string `db:"id" json:"id" validate:"required"`
	GroupID string `db:"group_id" json:"group_id"`
}

type ChatLine struct {
	ID        string    `db:"id" json:"id"`
	ChatID    string    `db:"chat_id" json:"chat_id" validate:"required"`
	UserID    string    `db:"user_id" json:"user_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	Message   string    `db:"message" json:"message"`
}

type ChatUser struct {
	ID     string `db:"id" json:"id"`
	UserID string `db:"user_id" json:"user_id" validate:"required"`
	ChatID string `db:"chat_id" json:"chat_id" validate:"required"`
}

func NewChatLine() *ChatLine {
	return &ChatLine{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
	}
}
