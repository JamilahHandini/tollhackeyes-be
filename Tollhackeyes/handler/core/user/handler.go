package user

import (
	"context"
	//"log"

	//"time"

	firebase "firebase.google.com/go"

	//"google.golang.org/api/option"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jamilahhandini/tollhackeyes/config"
	"github.com/jamilahhandini/tollhackeyes/models/dbscan"
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

	client, err := h.Database.Database(context.Background())
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	ref := client.NewRef("Drivers")
	
	// var s data.User
	// if err := ref.Get(context.Background(), &s); err != nil {
	// 	log.Fatalln("error in reading from firebase DB: ", err)
	//  }

	err = ref.Child(uuid.NewString()).Set(context.TODO(), dataReqs)
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	response := c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil",
	})

	return response
}

func (h UserHandler) GetInfoPerjalanans(c *fiber.Ctx) error{

	dataReqs := new(data.User)

	ctx := context.Background()
	err := c.BodyParser(dataReqs)
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	client, err := h.Database.Database(ctx)
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	ref := client.NewRef("Drivers/dvr1")
	// read from user_scores using ref
	s := &dbscan.InfoPerjalanan{}
	err = ref.Child("info_perjalanans").Get(context.Background(), s)
		if err != nil {
			c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}

	response := c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": s,
	})

	return response
}

func (h UserHandler) GetCurrentJarak(c *fiber.Ctx) error{

	dataReqs := new(data.User)

	ctx := context.Background()
	err := c.BodyParser(dataReqs)
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	client, err := h.Database.Database(ctx)
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	ref := client.NewRef("Drivers/dvr1")
	// read from user_scores using ref
	s := &dbscan.Perjalanan{}
	err = ref.Child("info_perjalanans").OrderByChild("").EqualTo(true).Get(ctx, s)
		if err != nil {
			c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}

	response := c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": s,
	})

	return response
}
