package handler

import (
	firebase "firebase.google.com/go"

	"github.com/jamilahhandini/tollhackeyes/config"
	"github.com/jamilahhandini/tollhackeyes/handler/core"
	"github.com/jamilahhandini/tollhackeyes/usecase"

)

//* Struct
type Handler struct {
	Core    core.CoreHandler
}

//* Handler Constructor
func NewHandler(uc usecase.Usecase, conf *config.Config, fb *firebase.App) Handler {
	return Handler{
		Core:    core.NewCoreHandler(uc, conf, fb),
	}
}
