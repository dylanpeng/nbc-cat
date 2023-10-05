package types

import "github.com/golang-jwt/jwt/v5"

type AdminClaims struct {
	jwt.RegisteredClaims
	UserId int64 `json:"user_id"`
}
