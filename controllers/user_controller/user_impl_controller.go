package user_controller

import (
	"errors"
	"go-jti/config"
	"go-jti/dto/response"
	"go-jti/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type userController struct {
}

func NewUserController() UserController {
	return &userController{}
}

func (c *userController) GoogleLogin(ctx *fiber.Ctx) error {
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")

	ctx.Status(fiber.StatusSeeOther)
	ctx.Redirect(url)
	return ctx.JSON(url)
}

func (c *userController) GoogleCallback(ctx *fiber.Ctx) error {
	state := ctx.Query("state")
	if state != "randomstate" {
		return ctx.SendString("States don't Match!!")
	}

	code := ctx.Query("code")

	if code == "" {
		return response.Error(ctx, "failed", http.StatusUnauthorized, errors.New("authorization code not provided!"))
	}

	tokenRes, err := utils.GetGoogleOauthToken(code)

	if err != nil {
		return response.Error(ctx, "failed", http.StatusBadGateway, err)
	}

	user, err := utils.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)
	if err != nil {
		return response.Error(ctx, "failed", http.StatusBadGateway, err)
	}

	token, err := utils.CreateToken(*user)
	if err != nil {
		return response.Error(ctx, "failed", http.StatusBadRequest, err)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: token,
	})

	return ctx.Redirect("/choose")

	// googlecon := config.GoogleConfig()

	// token, err := googlecon.Exchange(context.Background(), code)
	// if err != nil {
	// 	return ctx.SendString("Code-Token Exchange Failed")
	// }

	// resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	// if err != nil {
	// 	return ctx.SendString("User Data Fetch Failed")
	// }

	// userData, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return ctx.SendString("JSON Parsing Failed")
	// }

	// return ctx.SendString(string(userData))
}
