package helper

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TTL struct {
	minute int
	hour   int
	day    int
}

type Config struct {
	ttl    TTL
	secret []byte
}

type Helper struct {
	cfg *Config
}

func NewHelper(cfg *Config) (*Helper, error) {
	return &Helper{cfg}, nil
}

func (h *Helper) Authenticate(ctx context.Context, payload interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	b, _ := json.Marshal(payload)
	claims["payload"] = string(b)

	claims["exp"] = time.Now().Add(time.Minute*time.Duration(h.cfg.ttl.minute) + time.Hour*time.Duration(h.cfg.ttl.hour) + time.Hour*24*time.Duration(h.cfg.ttl.day)).Unix()

	return token.SignedString(h.cfg.secret)
}

func (h *Helper) ParseTokenString(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return h.cfg.secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		payload := claims["payload"]

		return payload, nil
	}

	return nil, errors.New("failed in parsing")
}
