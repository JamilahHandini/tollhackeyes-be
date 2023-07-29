package usecase

import (
	firebase "firebase.google.com/go"

	"github.com/jamilahhandini/tollhackeyes/config"
	"github.com/jamilahhandini/tollhackeyes/repository"
	"github.com/jamilahhandini/tollhackeyes/usecase/user"

)

//* Struct
type Usecase struct {
	User user.UserUsecase
}

//* Usecase Constructor
func NewUsecase(repo repository.Repository, conf *config.Config, fb *firebase.App) Usecase {
	return Usecase{
		User: user.NewUserUsecase(repo, conf, fb),
	}
}
