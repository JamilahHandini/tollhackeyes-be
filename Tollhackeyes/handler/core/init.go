package core

import (
	firebase "firebase.google.com/go"

	"github.com/jamilahhandini/tollhackeyes/config"
	"github.com/jamilahhandini/tollhackeyes/usecase"
	"github.com/jamilahhandini/tollhackeyes/handler/core/user"

)

//* Struct
type CoreHandler struct {
	User user.UserHandler
}

//* CoreHandler Constructor
func NewCoreHandler(uc usecase.Usecase, conf *config.Config, fb *firebase.App) CoreHandler {
	return CoreHandler{
		User: user.NewUserHandler(uc,conf,fb),
	}
}
