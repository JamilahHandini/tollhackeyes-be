package user

import (
	"context"
	firebase "firebase.google.com/go"

	"github.com/jamilahhandini/tollhackeyes/config"
	data "github.com/jamilahhandini/tollhackeyes/models/dbscan"
	"github.com/jamilahhandini/tollhackeyes/repository"
)

// * Interface
type Usecase interface {
	Register(ctx context.Context, user *data.User) (bool, string)
	
}

// * Struct
type UserUsecase struct {
	Repo repository.Repository
	Conf *config.Config
	Database *firebase.App
}

// * Constructor
func NewUserUsecase(repository repository.Repository, conf *config.Config, fb *firebase.App) UserUsecase {
	return UserUsecase{
		Repo: repository,
		Conf: conf,
		Database: fb,
	}
}

// * Tampilkan Semua Data Berat
func (u UserUsecase) Register(ctx context.Context, user *data.User) (bool, string) {

	//* Repo
	status, message := u.Repo.User.Register(ctx, user)
	if status == false {
		//* Error
		return status, message
	}

	//* Success
	return status, message
}
