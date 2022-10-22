package router

import (
	_ "github.com/ario-team/glassnode-api/docs"
	"github.com/ario-team/glassnode-api/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

var docRoute fiber.Router = web.Server.Group("/doc")

func InitializeDocRoute() {
	docRoute.Get("/*", swagger.HandlerDefault)
}
