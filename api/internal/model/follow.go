package models

type Follower struct {
	ID string `db:"id" json:"id"`
	SourceID string `db:"source_id" json:"source_id"`
	TargetID string `db:"target_id" json:"target_id"`
}