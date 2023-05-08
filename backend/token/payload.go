package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrExpiredToken = errors.New("token has expired")
var ErrInvalidToken = errors.New("token is invalid")

// Payload はやりとりに必要な属性情報（Claim）を持つ JSON 形式のデータ
type Payload struct {
	ID         uuid.UUID `json:"id"`
	EmployeeID string    `json:"employee_id"`
	IsAdmin    bool      `json:"is_admin"`
	IssuedAt   time.Time `json:"issued_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}

// NewPayload は有効期限を指定して ユーザーのメールアドレスから Payload を生成する
func NewPayload(employeeID string, isAdmin bool, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:         tokenID,
		EmployeeID: employeeID,
		IsAdmin:    isAdmin,
		IssuedAt:   time.Now(),
		ExpiredAt:  time.Now().Add(duration),
	}

	return payload, nil
}

// jwt.NewWithClaims の引数に渡すために、Payload に対して Valid() メソッドを実装する
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
