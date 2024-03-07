package jwttoken

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash"
	"social-network/pkg/errors"
	"strings"
	"time"
)

type Algorithm struct {
	hash      hash.Hash
	algorithm string
}

type Header struct {
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

func (a *Algorithm) sum(data []byte) []byte {
	return a.hash.Sum(data)
}

func (a *Algorithm) reset() {
	a.hash.Reset()
}

func (a *Algorithm) write(data []byte) (int, error) {
	return a.hash.Write(data)
}

func (a *Algorithm) NewHeader() *Header {
	return &Header{
		Typ: "JWT",
		Alg: a.algorithm,
	}
}

// Sign signs token
func (a *Algorithm) Sign(unsignedToken string) ([]byte, error) {
	_, err := a.write([]byte(unsignedToken))
	if err != nil {
		return nil, err
	}

	encodedToken := a.sum(nil)
	a.reset()

	return encodedToken, nil
}

// Encode returns signed JWT token
func (a *Algorithm) Encode(payload *Claims) (string, error) {
	header := a.NewHeader()

	jsonHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	b64Header := base64.RawURLEncoding.EncodeToString(jsonHeader)

	jsonPayload, err := json.Marshal(payload.claimsMap)
	if err != nil {
		return "", err
	}

	b64Payload := base64.RawURLEncoding.EncodeToString(jsonPayload)

	unsignedToken := b64Header + "." + b64Payload

	signature, err := a.Sign(unsignedToken)
	if err != nil {
		fmt.Println(err)
	}

	b64Signature := base64.RawURLEncoding.EncodeToString(signature)

	token := b64Header + "." + b64Payload + "." + b64Signature

	return token, nil
}

// Decode returns claims, does not validate them
func (a *Algorithm) Decode(token string) (*Claims, error) {
	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 3 {
		return nil, errors.New("invalid token format")
	}
	b64Payload := tokenParts[1]

	payload, err := base64.RawURLEncoding.DecodeString(b64Payload)
	if err != nil {
		return nil, err
	}

	var claims map[string]interface{}
	if err = json.Unmarshal(payload, &claims); err != nil {
		return nil, err
	}

	return &Claims{claimsMap: claims}, nil
}

// Validate verifies token validity, returns nil on success
func (a *Algorithm) Validate(token string) error {
	_, err := a.DecodeAndValidate(token)
	return err
}

// DecodeAndValidate returns token's claims and it's valid (nil if valid)
func (a *Algorithm) DecodeAndValidate(token string) (*Claims, error) {
	claims, err := a.Decode(token)
	if err != nil {
		return claims, err
	}

	if err = a.validateSignature(token); err != nil {

		err = errors.Join(errors.New("failed to validate signature"), err)
		return claims, err
	}

	if err = a.validateExp(claims); err != nil {
		err = errors.Join(errors.New("failed to validate exp"), err)
		return claims, err
	}

	return claims, nil
}

func (a *Algorithm) validateSignature(token string) error {
	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 3 {
		return errors.New("invalid token format")
	}
	b64Header := tokenParts[0]
	b64Payload := tokenParts[1]
	b64Signature := tokenParts[2]

	signedAttempt, err := a.Sign(b64Header + "." + b64Payload)
	if err != nil {
		return errors.Join(errors.New("failed to sign token for validation"), err)
	}

	b64SignedAttempt := base64.RawURLEncoding.EncodeToString(signedAttempt)

	if !hmac.Equal([]byte(b64Signature), []byte(b64SignedAttempt)) {
		return errors.New("invalid signature")
	}

	return nil
}

func (a *Algorithm) validateExp(claims *Claims) error {
	exp, err := claims.GetTime("exp")
	if err != nil {
		return err
	}
	if exp.Before(time.Now()) {
		return errors.New("token has expired")
	}

	return nil
}

func HmacSha256(key string) Algorithm {
	return Algorithm{
		hash:      hmac.New(sha256.New, []byte(key)),
		algorithm: "HS256",
	}
}
