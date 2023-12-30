package pages_controller

import "github.com/gofiber/fiber/v2"

type PagesController interface {
	IndexPages(ctx *fiber.Ctx) error
	ChoosePages(ctx *fiber.Ctx) error
	InputPages(ctx *fiber.Ctx) error
	OutputPages(ctx *fiber.Ctx) error
}
