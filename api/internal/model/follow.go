package model

type Follower struct {
	ID       string `db:"id" json:"id" validate:"required"`
	SourceID string `db:"source_id" json:"source_id" validate:"required"`
	TargetID string `db:"target_id" json:"target_id" validate:"required"`
}
