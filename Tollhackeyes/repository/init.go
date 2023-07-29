package repository

import (
	firebase "firebase.google.com/go"
	"github.com/jamilahhandini/tollhackeyes/repository/user"
	"github.com/jamilahhandini/tollhackeyes/config"
)

//* Struct
type Repository struct {
	User user.UserRepo
}

//* Repository Constructor
func NewRepository(conf *config.Config, fb *firebase.App) Repository {
	return Repository{
		User: user.NewUserRepo(conf,fb),
	}
}