package model

type Follow struct {
	SourceId string `json:"source_id"`
	TargetId string `json:"target_id"`
}

const (
	PUBLIC   = 1
	PRIVATE  = 2
	SELECTED = 3
)
