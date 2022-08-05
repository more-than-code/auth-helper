package helper

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Config struct {
	TtlMinute int
	TtlHour   int
	TtlDay    int
	Secret    []byte
}

type Helper struct {
	cfg *Config
}

func NewHelper(cfg *Config) (*Helper, error) {
	return &Helper{cfg}, nil
}

func (h *Helper) Authenticate(str string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["str"] = str

	claims["exp"] = time.Now().Add(time.Minute*time.Duration(h.cfg.TtlMinute) + time.Hour*time.Duration(h.cfg.TtlHour) + time.Hour*24*time.Duration(h.cfg.TtlDay)).Unix()

	return token.SignedString(h.cfg.Secret)
}

func (h *Helper) ParseTokenString(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return h.cfg.Secret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		str := claims["str"]

		return str.(string), nil
	}

	return "", errors.New("failed in parsing")
}
