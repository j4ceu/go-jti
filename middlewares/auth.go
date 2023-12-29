package middlewares

import (
	"errors"
	"go-jti/dto/response"
	"go-jti/utils"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Auth() fiber.Handler {
	return func(context *fiber.Ctx) error {

		authorizationToken := context.Get("Authorization")
		tokenString := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

		if !strings.Contains(authorizationToken, "Bearer") {
			baseResponse := response.Error(context, "failed", http.StatusUnauthorized, errors.New("request does not contain an access token"))
			return baseResponse
		}

		if tokenString == "" {
			baseResponse := response.Error(context, "failed", http.StatusUnauthorized, errors.New("request does not contain an access token"))
			return baseResponse
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			baseResponse := response.Error(context, "failed", http.StatusUnauthorized, err)
			return baseResponse
		}

		context.Locals("claims", claims)

		return context.Next()
	}
}
