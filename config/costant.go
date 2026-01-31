package config

import "time"

const (
	JwtSignKey           = "jwt-secret"
	AccessTokenSubject   = "ac"
	RefreshTokenSubject  = "rt"
	AccessTokenExpireDuration = time.Hour*24
	RefreshTokenExpireDuration=time.Hour*24
	AuthMiddlewareContextKey="claims"
)