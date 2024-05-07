package model

import "errors"

type Type struct {
	Privacy PrivacyType
	Request RequestType
	Member  MemberType
	Event   EventType
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

type EventType struct {
	Going      int
	NotGoing   int
	Interested int
	Maybe      int
	Values     map[string]int
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
		Event: EventType{
			Going:      1,
			NotGoing:   2,
			Interested: 3,
			Maybe:      4,
			Values: map[string]int{
				"going":      1,
				"notgoing":   2,
				"interested": 3,
				"maybe":      4,
			},
		},
	}
}

func (t *Type) IntToString(num int) (string, error) {
	for key, val := range t.Privacy.Values {
		if val == num {
			return key, nil
		}
	}
	for key, val := range t.Request.Values {
		if val == num {
			return key, nil
		}
	}
	for key, val := range t.Member.Values {
		if val == num {
			return key, nil
		}
	}
	for key, val := range t.Event.Values {
		if val == num {
			return key, nil
		}
	}
	return "", errors.New("no valid num")
}
