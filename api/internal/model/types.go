package model

type Type struct {
	Privacy PrivacyType
	Request RequestType
	Member  MemberType
}

type PrivacyType struct {
	Public   int
	Private  int
	Selected int
	Values   map[string]int
}

type RequestType struct {
	Notification int
	Follow       int
	Invite       int
	Values       map[string]int
}

type MemberType struct {
	Admin  int
	User   int
	Values map[string]int
}

func InitializeTypes() Type {
	return Type{
		Privacy: PrivacyType{
			Public:   1,
			Private:  2,
			Selected: 3,
			Values: map[string]int{
				"public":   1,
				"private":  2,
				"selected": 3,
			},
		},
		Request: RequestType{
			Notification: 1,
			Follow:       2,
			Invite:       3,
			Values: map[string]int{
				"notification": 1,
				"follow":       2,
				"invite":       3,
			},
		},
		Member: MemberType{
			Admin: 1,
			User:  2,
			Values: map[string]int{
				"admin": 1,
				"user":  2,
			},
		},
	}
}
