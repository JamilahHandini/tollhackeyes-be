package user

import (
	"context"

	firebase "firebase.google.com/go"

	"github.com/jamilahhandini/tollhackeyes/config"
	data "github.com/jamilahhandini/tollhackeyes/models/dbscan"
)

// * Interface
type Repository interface {
	Register(ctx context.Context, user *data.User) (bool, string)
	
}

// * Data Struct
type UserRepo struct {
	Conf *config.Config
	Database *firebase.App
}

// * Constructor
func NewUserRepo(conf *config.Config, fb *firebase.App) UserRepo {
	return UserRepo{
		Conf: conf,
		Database: fb,
	}
}

func (r UserRepo) Register(ctx context.Context, user *data.User) (bool, string) {
	
	client, err := r.Database.Database(ctx)
	if err != nil {
		return false, "client"
	}

	ref := client.NewRef("Drivers")
	
	_, err = ref.Push(ctx, user)
	if err != nil {
		return false, "Message: Register gagal"
	}

	return true, "Message: Register berhasil"
}


