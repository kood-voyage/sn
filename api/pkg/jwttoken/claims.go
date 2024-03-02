package jwttoken

import (
	"fmt"
	"social-network/pkg/errors"
	"time"
)

type Claims struct {
	claimsMap map[string]interface{}
}

func NewClaims() *Claims {
	claimsMap := make(map[string]interface{})

	claims := &Claims{claimsMap: claimsMap}

	claims.SetTime("iat", time.Now())
	return claims
}

func (c *Claims) Set(key string, value interface{}) {
	c.claimsMap[key] = value
}

func (c *Claims) Get(key string) (interface{}, error) {
	result, ok := c.claimsMap[key]
	if !ok {
		fmt.Println("test123")
		return "", fmt.Errorf("claim (%s) doesn't exist", key)		
	}

	return result, nil
}

func (c *Claims) SetTime(key string, value time.Time) {
	c.Set(key, value.Unix())
}

func (c *Claims) GetTime(key string) (time.Time, error) {
	raw, err := c.Get(key)
	if err != nil {
		return time.Unix(0, 0), err
	}

	timeFloat, ok := raw.(float64)
	if !ok {
		return time.Unix(0, 0), errors.Join(err, fmt.Errorf("claim isn't a valid float"))
	}

	return time.Unix(int64(timeFloat), 0), nil
}

func (c *Claims) HasClaim(key string) bool {
	_, ok := c.claimsMap[key]
	return ok
}
