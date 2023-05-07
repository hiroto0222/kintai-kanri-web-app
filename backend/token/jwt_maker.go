package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const minSecretKeySize = 32

// JWTMaker は JWT を生成するための構造体
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker は JWTMaker を生成する
func NewJWTMaker(secretKey string) (Maker, error) {
	// JWTの秘密鍵が強度なランダムキーを生成するためには > 32
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size, must be atleast %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// CreateToken は JWT トークンをユーザーのメールアドレスで生成する
func (maker *JWTMaker) CreateToken(email string, duration time.Duration) (string, error) {
	payload, err := NewPayload(email, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

// VerifyToken は JWT トークンの署名を検証し、そのトークンが有効であるかどうかを確認する
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	//
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// SigningMethodHS256を使用しているため
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		// jwt.ParseWithClaims は、トークンが有効期限切れの場合にはerr.InnerにPayload.Valid()のエラー情報を返す
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
