package initialize

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {

	// Run server.
	if err := a.Listen(":8000"); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}

func FiberMiddleware(a *fiber.App) {
	a.Use(
		//Middleware
		cors.New(),
		recover.New(),
	)
}
