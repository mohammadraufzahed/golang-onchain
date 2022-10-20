package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var Server *fiber.App = fiber.New()

// @title        Glassnode Service
// @version      1
// @description  Glassnode clone for use in Iran
// @contact.name ArioDev
// @contact.url  https://ariodev.com
// @host         localhost:3000
// @BasePath     /
func Start() {

	Server.Listen(":3000")
}

func InitailizeMiddlewares() {
	Server.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://127.0.0.1:3000",
	}))
}
