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

	ref := client.NewRef("Drivers/dvr1/info_perjalanans")
	// read from user_scores using ref

	data, errData := ref.OrderByChild("status_perjalanan").EqualTo(true).GetOrdered(ctx)
	
	if errData != nil {
		c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	var key string 
	for _, value := range data {
		key = value.Key()
	}

	ref = client.NewRef("Drivers/dvr1/info_perjalanans/")
	s := &dbscan.Perjalanan{}
	err = ref.Child(key).Get(context.Background(), s)
		if err != nil {
			c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}

	response := c.Status(400).JSON(fiber.Map{
		"success": true,
		"message": s,
	})

	return response
}

func (h UserHandler) TopUp(c *fiber.Ctx) error{

	dataReqs := new(data.EToll)

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
	

	s := &dbscan.EToll{}
	err = ref.Child("e_tolls").Get(context.Background(), s)
		if err != nil {
			c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}

	saldo := 0
	saldo = s.Saldo + dataReqs.Saldo
	
	ref2 := client.NewRef("Drivers").Child("dvr1").Child("e_tolls")
	err = ref2.Update(ctx, map[string]interface{}{"saldo": saldo })
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	response := c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "berhasil topup",
	})

	return response
}


func (h UserHandler) ClaimReward(c *fiber.Ctx) error{

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

	ref := client.NewRef("Drivers/dvr1/info_perjalanans")
	// read from user_scores using ref

	data, errData := ref.OrderByChild("status_perjalanan").EqualTo(true).GetOrdered(ctx)
	
	if errData != nil {
		c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	var key string 
	for _, value := range data {
		key = value.Key()
	}

	ref = client.NewRef("Drivers/dvr1/info_perjalanans/")
	s := &dbscan.Perjalanan{}
	err = ref.Child(key).Get(context.Background(), s)
		if err != nil {
			c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}

	if (s.JumlahPenumpang == 4){
		ref2 := client.NewRef("Drivers").Child("dvr1").Child("info_perjalanans").Child(key)
		err = ref2.Update(ctx, map[string]interface{}{"reward":  15})
		if err != nil {
			c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}
	}else if(s.JumlahPenumpang == 3){
		ref2 := client.NewRef("Drivers").Child("dvr1").Child("info_perjalanans").Child(key)
		err = ref2.Update(ctx, map[string]interface{}{"reward":  10})
		if err != nil {
			c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}
	}else if(s.JumlahPenumpang == 2){
		ref2 := client.NewRef("Drivers").Child("dvr1").Child("info_perjalanans").Child(key)
		err = ref2.Update(ctx, map[string]interface{}{"reward":  5})
		if err != nil {
			c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}
	}else if(s.JumlahPenumpang == 1){
		ref2 := client.NewRef("Drivers").Child("dvr1").Child("info_perjalanans").Child(key)
		err = ref2.Update(ctx, map[string]interface{}{"reward":  0})
		if err != nil {
			c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}
	}

	response := c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil Claim",
	})

	return response
}


func (h UserHandler) Pay(c *fiber.Ctx) error{

	dataReqs := new(data.Pay)

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

	ref := client.NewRef("Drivers/dvr1/info_perjalanans")
	// read from user_scores using ref

	data, errData := ref.OrderByChild("status_perjalanan").EqualTo(true).GetOrdered(ctx)
	
	if errData != nil {
		c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	var key string 
	for _, value := range data {
		key = value.Key()
	}

	ref = client.NewRef("Drivers/dvr1/info_perjalanans/")
	s := &dbscan.Perjalanan{}
	err = ref.Child(key).Get(context.Background(), s)
		if err != nil {
			c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}
	
	var total_pay int
	var persen float64
	var diskon int
	if(s.Reward != 0){ 
		persen = float64(s.Reward)/ float64(100)
		a := float64(dataReqs.Tarif) * persen
		diskon = int(a)
		total_pay = dataReqs.Tarif - diskon
	}

	ref3 := client.NewRef("Drivers/dvr1")
	// read from user_scores using ref
	

	s1 := &dbscan.EToll{}
	err = ref3.Child("e_tolls").Get(context.Background(), s1)
		if err != nil {
			c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}

	saldo := 0
	if(s1.Saldo > total_pay){
		saldo = s1.Saldo - total_pay
		ref2 := client.NewRef("Drivers").Child("dvr1").Child("e_tolls")
		err = ref2.Update(ctx, map[string]interface{}{"saldo":  saldo})
		if err != nil {
			c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}
	}else{
		c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Saldo anda Kurang",
		})
	}
	

	response := c.Status(200).JSON(fiber.Map{
		"success": true,
		"tarif_asli": dataReqs.Tarif,
		"tarif_reward" : total_pay,
	})

	return response
}
