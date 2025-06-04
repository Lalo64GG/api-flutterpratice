package middlewares

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	Id int64  `json:"id"`
	jwt.RegisteredClaims
}