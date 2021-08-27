package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type JWTAuthServiceAdapter struct {
	service JWTService
}

func NewJwtAuthServiceAdabter() JWTAuthServiceAdapter {
	return JWTAuthServiceAdapter{
		service: JWTAuthService(),
	}
}

func (adabter JWTAuthServiceAdapter) GenerateToken (key string) string  {
	return adabter.service.GenerateToken(key , true)
}

func (adabter JWTAuthServiceAdapter) ValidateToken (token string) error  {
	_, err := adabter.service.ValidateToken(token)

	return err
}

func (adabter JWTAuthServiceAdapter) Authorize (token string) (map[string]interface{} , error)  {
	validatedToken, err := JWTAuthService().ValidateToken(token)

	if err != nil {
		return nil, err
	}

	if validatedToken.Valid {
		claims := validatedToken.Claims.(jwt.MapClaims)
		return claims , nil
	} else {
		return nil , errors.New("unauthorized")
	}
}
