package web

import "github.com/gofiber/fiber/v2"

var Server *fiber.App = fiber.New()

// @title Glassnode Service
// @version 1
// @description Glassnode clone for use in Iran
// @contact.name ArioDev
// @contact.url https://ariodev.com
// @host localhost:3000
// @BasePath /
func Start() {
	Server.Listen(":3000")
}
