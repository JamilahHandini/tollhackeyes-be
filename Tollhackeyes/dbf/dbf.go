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
		DatabaseURL: conf.Database.DatabaseURL,
	}
	// replace this with your service account file location
	opt := option.WithCredentialsFile(conf.Database.Path) 
	app, err := firebase.NewApp(ctx, configFirebase, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	return app
  }
  