package auth

import (
	"auth_micro/config/auth"
	"auth_micro/internal/services/auth/jwt"
)

type AuthServiceAdabter interface {
	GenerateToken(key string) string
	ValidateToken(token string) error
	Authorize(token string) (map[string]interface{} , error)
}


func GetAdabter() AuthServiceAdabter {
	switch auth.DefaultAuthService {
		case "jwt":
			return jwt.NewJwtAuthServiceAdabter()
	}

	return nil
}