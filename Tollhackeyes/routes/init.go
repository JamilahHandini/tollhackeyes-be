package routes

import (
	"github.com/jamilahhandini/tollhackeyes/config"
	"github.com/jamilahhandini/tollhackeyes/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h handler.Handler, conf *config.Config) {
	
	//* Api v1
	api := app.Group("/api/tollhackeyes")

	//* Core endpoint
	RouteUser(api.Group("/content"), conf, h)

}
