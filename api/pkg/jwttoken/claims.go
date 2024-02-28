package jwttoken

import "time"

type Claims struct {
	UserID string `json:"user_id"`
	Iat    int64  `json:"iat"`
	Exp    int64  `json:"exp"`
}

func NewClaims(id string, exp int64) *Claims {
	return &Claims{
		UserID: id,
		Iat:    time.Now().Unix(),
		Exp:    exp,
	}
}
