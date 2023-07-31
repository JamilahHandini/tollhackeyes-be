package dbf

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/jamilahhandini/tollhackeyes/config"
	"google.golang.org/api/option"
)
  
  func NewConnection(conf *config.Config) *firebase.App{
	ctx := context.Background()
	configFirebase := &firebase.Config{
		DatabaseURL: "https://toll-wise-driver-default-rtdb.firebaseio.com/",
	}
	opt := option.WithCredentialsFile("toll-wise-driver-firebase-adminsdk-2yj8r-d28e54d76b.json")
	app, err := firebase.NewApp(ctx, configFirebase, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	return app
  }
  