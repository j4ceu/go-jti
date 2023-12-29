package routes

import (
	"go-jti/initialize"
	"go-jti/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	initialize.FiberMiddleware(app)

	app.Get("/google_login", initialize.UserController.GoogleLogin)
	app.Get("/google_callback", initialize.UserController.GoogleCallback)

	v1 := app.Group("/v1")

	phoneNumber := v1.Group("/phone-number")
	phoneNumber.Post("", initialize.PhoneNumberController.CreatePhoneNumber)
	phoneNumber.Put("/:id", initialize.PhoneNumberController.UpdatePhoneNumber)
	phoneNumber.Delete("/:id", initialize.PhoneNumberController.DeletePhoneNumber)
	phoneNumber.Get("/even", middlewares.Auth(), initialize.PhoneNumberController.FindEvenPhoneNumber)
	phoneNumber.Get("/odd", initialize.PhoneNumberController.FindOddPhoneNumber)
	phoneNumber.Get("/:id", initialize.PhoneNumberController.FindPhoneNumberByID)
}
