package token

import "time"

// Maker はトークン（JWT・Paseto）を生成するインターフェース
type Maker interface {
	// CreateToken はトークンをユーザーのメールアドレスで生成する
	CreateToken(email string, isAdmin bool, duration time.Duration) (string, *Payload, error)
	// VerifyToken はトークンの署名を検証し、そのトークンが有効であるかどうかを確認する
	VerifyToken(token string) (*Payload, error)
}
