package pages_controller

import (
	"go-jti/utils"

	"github.com/gofiber/fiber/v2"
)

type pagesController struct {
}

func NewPagesController() PagesController {
	return &pagesController{}
}

func (c *pagesController) IndexPages(ctx *fiber.Ctx) error {
	return ctx.SendFile("utils/pages/index.html")
}

func (c *pagesController) ChoosePages(ctx *fiber.Ctx) error {
	token := ctx.Cookies("token")
	claims, err := utils.ValidateToken(token)
	if err != nil {
		return ctx.SendFile("utils/pages/not-logged-in.html")
	}

	if token != "" {
		return ctx.Render("utils/pages/choose.html", fiber.Map{
			"Name": claims.Name,
		})
	} else {
		return ctx.SendFile("utils/pages/not-logged-in.html")
	}

}

func (c *pagesController) InputPages(ctx *fiber.Ctx) error {
	token := ctx.Cookies("token")
	if token != "" {
		return ctx.SendFile("utils/pages/input.html")
	} else {
		return ctx.SendFile("utils/pages/not-logged-in.html")
	}
}

func (c *pagesController) OutputPages(ctx *fiber.Ctx) error {
	token := ctx.Cookies("token")
	if token != "" {
		return ctx.SendFile("utils/pages/output.html")
	} else {
		return ctx.SendFile("utils/pages/not-logged-in.html")
	}
}
