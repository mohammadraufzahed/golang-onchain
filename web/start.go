package web

import "github.com/gofiber/fiber/v2"

var Server *fiber.App = fiber.New()

func Start() {
	Server.Listen(":3000")
}
