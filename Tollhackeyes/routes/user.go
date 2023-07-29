package routes

import (
	"github.com/jamilahhandini/tollhackeyes/config"
	"github.com/jamilahhandini/tollhackeyes/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteUser(api fiber.Router, conf *config.Config, handler handler.Handler) {

	api.Post("/register", handler.Core.User.Register)
	//api.Get("/getcontentreviewsummary", middleware.JWTValidator(conf, appLoger), handler.Core.Content.GetContentReview)
	
}
