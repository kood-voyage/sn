package model

type User struct {
	ID string `db:"id" json:"id" validate:"required"`
	MemberType int `json:"member_type"`
}
