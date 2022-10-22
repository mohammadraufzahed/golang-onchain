package router

import (
	"fmt"
	"time"

	"github.com/ario-team/glassnode-api/web"
	"github.com/gofiber/fiber/v2"
)

func InitializeRouter() {
	InitializeDocRoute()
	InitializeTopGroupRouter()
	InitializeMiddleGroupRouter()
	InitializeEndpointRouter()
	web.Server.Get("/", func(c *fiber.Ctx) error {
		go gofiber()
		return c.SendString("Started")
	})
}

func gofiber() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		time.Sleep(time.Second * 2)
	}
}
