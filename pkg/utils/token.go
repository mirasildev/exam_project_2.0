package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mirasildev/exam_project_2.0/config"
)


type TokenParams struct {
	UserID int64
	UserName string
	Email string
	UserType string
	Duration time.Duration
}


func CreateToken(cfg *config.Config, params *TokenParams) (string, *Payload, error) {
	payload, err := NewPayload(params)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(cfg.AuthSecretKey))
	return token, payload, err
}


func VerifyToken(cfg *config.Config, token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			 return nil, ErrInvalidToken
		}
		return []byte(cfg.AuthSecretKey), nil
	}

	jwtToken ,err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}