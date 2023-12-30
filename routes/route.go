package routes

import (
	"go-jti/initialize"
	"go-jti/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	initialize.FiberMiddleware(app)

	app.Static("/", "utils/pages")
	app.Get("/", initialize.PagesController.IndexPages)
	app.Get("/choose", initialize.PagesController.ChoosePages)
	app.Get("/input", initialize.PagesController.InputPages)
	app.Get("/output", initialize.PagesController.OutputPages)

	app.Get("/google_login", initialize.UserController.GoogleLogin)
	app.Get("/google_callback", initialize.UserController.GoogleCallback)

	v1 := app.Group("/v1")

	phoneNumber := v1.Group("/phone-number")
	phoneNumber.Post("", middlewares.Auth(), initialize.PhoneNumberController.CreatePhoneNumber)
	phoneNumber.Put("/:id", middlewares.Auth(), initialize.PhoneNumberController.UpdatePhoneNumber)
	phoneNumber.Delete("/:id", middlewares.Auth(), initialize.PhoneNumberController.DeletePhoneNumber)
	phoneNumber.Get("/even", middlewares.Auth(), initialize.PhoneNumberController.FindEvenPhoneNumber)
	phoneNumber.Get("/odd", middlewares.Auth(), initialize.PhoneNumberController.FindOddPhoneNumber)
	phoneNumber.Get("find/:id", middlewares.Auth(), initialize.PhoneNumberController.FindPhoneNumberByID)
	phoneNumber.Get("/generate-number", middlewares.Auth(), initialize.PhoneNumberController.GeneratePhoneNumber)
}
