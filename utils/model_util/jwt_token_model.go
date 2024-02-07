package model_util

import "github.com/golang-jwt/jwt/v5"

type JwtTokenClaims struct {
	jwt.RegisteredClaims
	// apa aja yg akan dibuat di payload
	UserId   string   `json:"userId"`
	Role     string   `json:"role"`
	Services []string `json:"services"`
}
