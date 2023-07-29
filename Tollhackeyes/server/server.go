package server

import (
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jamilahhandini/tollhackeyes/config"
	"github.com/jamilahhandini/tollhackeyes/handler"
	"github.com/jamilahhandini/tollhackeyes/repository"
	"github.com/jamilahhandini/tollhackeyes/routes"
	"github.com/jamilahhandini/tollhackeyes/usecase"
)

// * Running Program
func Run(conf *config.Config, fb *firebase.App) {

	//* Initial Engine
	engine := html.New("./views", ".html")

	//* Initial Fiber App
	app := fiber.New(fiber.Config{
		AppName:      "Tollhackeyes",
		ServerHeader: "Go Fiber",
		Views:        engine,
		BodyLimit:    20 * 1024 * 1024,
	})
	
	
	repo := repository.NewRepository(conf, fb)
	usecase := usecase.NewUsecase(repo, conf, fb)
	handler := handler.NewHandler(usecase, conf, fb)

	routes.SetupRoutes(app, handler, conf)

	port := "8000"
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
