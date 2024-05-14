package model

import "time"

type Session struct {
	AcessID   string    `db:"access_id" json:"access_id"`
	UserID    string    `db:"user_id" json:"user_id"`
	CreatedAT time.Time `db:"timestamp" json:"timestamp"`
}
