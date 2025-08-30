package user

import (
	"github.com/Minhajxdd/Ephemr/internal/user/controller"
	repositories "github.com/Minhajxdd/Ephemr/internal/user/repository"
	"github.com/Minhajxdd/Ephemr/internal/user/service"
	"gorm.io/gorm"
)

type Module struct {
	UserController *controller.UserController
	UserRepository *repositories.UserRepository
}

func NewModule(db *gorm.DB) *Module {
	repo := repositories.NewUserRepository(db)
	srv := service.NewUserService(&repo)
	ctrl := controller.NewUserService(&srv)

	return &Module{
		UserController: &ctrl,
		UserRepository: &repo,
	}
}
