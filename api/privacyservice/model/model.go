package model

type PrivacyReq struct {
	ParentId string `json:"parent_id"`
	Privacy  int32  `json:"privacy"`
}
