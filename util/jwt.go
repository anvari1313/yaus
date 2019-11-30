package util

import (
	"github.com/anvari1313/yaus/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTGenerator interface {
	Generate() (string, error)
}

type jwtGenerator struct {
	ttl time.Duration
	key []byte
}

func CreateJWTGenerator(c config.JWT) *jwtGenerator {
	return &jwtGenerator{
		ttl: c.TTL,
		key: []byte(c.Key),
	}
}

func (j *jwtGenerator) Generate() (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(j.ttl).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString(j.key)
	if err != nil {
		return "", err
	}

	return t, nil
}
