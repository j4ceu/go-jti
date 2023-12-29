package phone_number_controller

import "github.com/gofiber/fiber/v2"

type PhoneNumberController interface {
	CreatePhoneNumber(ctx *fiber.Ctx) error
	UpdatePhoneNumber(ctx *fiber.Ctx) error
	DeletePhoneNumber(ctx *fiber.Ctx) error
	FindPhoneNumberByID(ctx *fiber.Ctx) error
	FindOddPhoneNumber(ctx *fiber.Ctx) error
	FindEvenPhoneNumber(ctx *fiber.Ctx) error
}
