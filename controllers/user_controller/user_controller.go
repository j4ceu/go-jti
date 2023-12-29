package user_controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	GoogleLogin(ctx *fiber.Ctx) error
	GoogleCallback(ctx *fiber.Ctx) error
}
