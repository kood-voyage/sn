package model

import "time"

type RequestReq struct {
	Id        string    `json:"id"`
	TypeId    int32     `json:"type_id"`
	SourceId  string    `json:"source_id"`
	TargetId  string    `json:"target_id"`
	ParentId  string    `json:"parent_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type RequestReqs struct {
	Requests []RequestReq `json:"requests"`
}

type RequestUserId struct {
	Id string `json:"user_id"`
}

const (
	NOTIFICATION = iota + 1
	FOLLOW
	INVITE
)
