package main

import (
	"go-jti/initialize"
	"go-jti/routes"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

func main() {
	initialize.Setup()

	app := fiber.New(fiber.Config{
		UnescapePath: true,
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
	})

	routes.Setup(app)

	initialize.StartServer(app)
}
