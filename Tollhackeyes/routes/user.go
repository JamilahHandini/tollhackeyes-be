package routes

import (
	"github.com/jamilahhandini/tollhackeyes/config"
	"github.com/jamilahhandini/tollhackeyes/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteUser(api fiber.Router, conf *config.Config, handler handler.Handler) {

	//api.Post("/register", handler.Core.User.Register)
	api.Post("/topup", handler.Core.User.TopUp)
	api.Post("/pay", handler.Core.User.Pay)
	api.Get("/claimreward", handler.Core.User.ClaimReward)
	api.Get("/info_perjalanans", handler.Core.User.GetInfoPerjalanans)
	api.Get("/current_perjalanan", handler.Core.User.GetCurrentJarak)
	
	
}
