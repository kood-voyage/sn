package model

type Type struct {
	Privacy PrivacyType
	Request RequestType
	Member  MemberType
	Image   ImageType
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

type ImageType struct {
	Banner  int
	Avatar  int
	Profile int
	Header  int
	Values  map[string]int
}

func InitializeTypes() Type {
	return Type{
		Privacy: PrivacyType{
			Public:   0,
			Private:  1,
			Selected: 2,
			Values: map[string]int{
				"public":   0,
				"private":  1,
				"selected": 2,
			},
		},
		Request: RequestType{
			Notification: 0,
			Follow:       1,
			Invite:       2,
			Values: map[string]int{
				"notification": 0,
				"follow":       1,
				"invite":       2,
			},
		},
		Member: MemberType{
			Admin: 0,
			User:  1,
			Values: map[string]int{
				"admin": 0,
				"user":  1,
			},
		},
		Image: ImageType{
			Banner:  0,
			Avatar:  1,
			Profile: 2,
			Header:  3,
			Values: map[string]int{
				"banner":   0,
				"avatar":  1,
				"profile": 2,
				"header": 3,
			},
		},
	}
}
