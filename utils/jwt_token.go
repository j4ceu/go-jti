package utils

import (
	"errors"
	"go-jti/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type JWTClaim struct {
	Name  string `json:"name"`
	Email string `json:"email"`

	jwt.StandardClaims
}

func CreateToken(user GoogleUserResult) (tokenString string, err error) {
	authConfig := config.LoadAuthConfig()
	claims := &JWTClaim{
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: authConfig.ExpHours.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(authConfig.Secret))
	return tokenString, err
}

func ValidateToken(signedToken string) (*JWTClaim, error) {
	authConfig := config.LoadAuthConfig()
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(authConfig.Secret), nil
		},
	)

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return claims, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return claims, err
	}
	return claims, err
}

func GetTokenClaims(ctx *fiber.Ctx) *JWTClaim {
	claims := JWTClaim{}

	if ctx.Locals("claims") != nil {
		claims = *ctx.Locals("claims").(*JWTClaim)
	}

	return &claims

}
