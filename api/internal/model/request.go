package models

import "time"

type Request struct {
	ID        string    `db:"id" json:"id"`
	TypeID    string    `db:"type_id" json:"type_id"`
	SourceID  string    `db:"source_id" json:"source_id"`
	TargetID  string    `db:"target_id" json:"target_id"`
	Message   string    `db:"message" json:"message"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type RequestType struct {
	ID          string `db:"id" json:"id"`
	Description string `db:"description" json:"description"`
}
