package user

import (
	"context"
	//"time"

	firebase "firebase.google.com/go"

	"github.com/gofiber/fiber/v2"
	"github.com/jamilahhandini/tollhackeyes/config"
	data "github.com/jamilahhandini/tollhackeyes/models/dbscan"
	"github.com/jamilahhandini/tollhackeyes/usecase"
	//"github.com/jamilahhandini/tollhackeyes/utils"
)

// * Struct
type UserHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Database *firebase.App
}

// * Constructor
func NewUserHandler(uc usecase.Usecase, conf *config.Config, fb *firebase.App) UserHandler {
	return UserHandler{
		Usecase: uc,
		Conf:    conf,
		Database:    fb,
	}
}

func (h UserHandler) Register(c *fiber.Ctx) error{

	dataReqs := new(data.User)
	err := c.BodyParser(dataReqs)
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	status,message := h.Usecase.User.Register(context.Background(), dataReqs)

	response := c.Status(200).JSON(fiber.Map{
		"success": status,
		"message": message,
	})

	return response
}
