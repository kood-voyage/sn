package jwttoken

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

var key = "this is key"
var algorithm = HmacSha256(key)

func RunTest(t *testing.T, command func(Algorithm)) {
	command(algorithm)
}

func TestEncodeAndValidateToken(t *testing.T) {
	RunTest(t, func(a Algorithm) {
		id := "123456"
		exp := time.Now().Add(time.Duration(100) * time.Hour)
		payload := NewClaims()
		payload.Set("id", id)
		payload.SetTime("exp", exp)

		token, err := algorithm.Encode(payload)
		if err != nil {
			t.Fatal(err)
		}

		err = algorithm.Validate(token)
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestDecodeInvalidToken(t *testing.T) {
	RunTest(t, func(a Algorithm) {
		tests := []struct {
			name  string
			token string
			want  string
		}{
			{
				name:  "empty token",
				token: "",
				want:  "invalid token format",
			},
			{
				name:  "invalid format",
				token: "abc",
				want:  "invalid token format",
			},
			{
				name:  "invalid format",
				token: "y04-387uh4.345h63h645",
				want:  "invalid token format",
			},
		}

		for _, test := range tests {
			_, got := a.Decode(test.token)
			if !reflect.DeepEqual(test.want, got.Error()) {
				t.Fatalf("expected: %v, got: %v", test.want, got.Error())
			}
		}
	})
}

func TestValidateToken(t *testing.T) {
	RunTest(t, func(a Algorithm) {
		id := "123456"
		exp := time.Now().Add(time.Duration(100) * time.Hour).Unix()
		payload := NewClaims()
		payload.Set("id", id)
		payload.Set("exp", exp)

		token, err := algorithm.Encode(payload)
		if err != nil {
			t.Fatal(err)
		}

		tokenParts := strings.Split(token, ".")
		invalidSignature := "q253wrw235egstf3rxjh54w76524q6tzgf"
		invalidToken := tokenParts[0] + "." + tokenParts[1] + "." + invalidSignature

		err = algorithm.Validate(invalidToken)
		if err == nil {
			t.Fatal(err)
		}
	})
}

func TestValidateExpiredToken(t *testing.T) {
	RunTest(t, func(a Algorithm) {
		id := "123456"
		exp := time.Now().Add(time.Duration(-1) * time.Hour)
		payload := NewClaims()
		payload.Set("id", id)
		payload.SetTime("exp", exp)

		token, err := algorithm.Encode(payload)
		if err != nil {
			t.Fatal(err)
		}

		err = algorithm.Validate(token)
		if err == nil {
			t.Fatal(err)
		}
	})
}

func TestValidateExternalToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiSm9obiBEb2UiLCJpYXQiOjE1MTY1MzkwMjIsImV4cCI6MzIzNDU3ODk4N30.POYGHYMlxbvgoT5LsGQZNxU9TawZYK9UBCqW3ILzV7w"
	alg := HmacSha256("this is secret key")
	err := alg.Validate(token)
	if err != nil {
		t.Fatal(err)
	}
}
