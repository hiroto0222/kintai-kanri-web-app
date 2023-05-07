package token

import "time"

// Maker はトークン（JWT・Paseto）を生成するインターフェース
type Maker interface {
	// CreateToken はトークンをユーザーのメールアドレスで生成する
	CreateToken(email string, duration time.Duration) (string, error)
	// VerifyToken はトークンの署名を検証し、そのトークンが有効であるかどうかを確認する
	VerifyToken(token string) (*Payload, error)
}
