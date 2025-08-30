package user

import (
	"github.com/Minhajxdd/Ephemr/internal/user/controller"
	repositories "github.com/Minhajxdd/Ephemr/internal/user/repository"
	"github.com/Minhajxdd/Ephemr/internal/user/service"
	"gorm.io/gorm"
)

type Module struct {
	Controller *controller.UserController
}

func NewModule(db *gorm.DB) *Module {
	repo := repositories.NewUserRepository(db)
	srv := service.NewUserService(&repo)
	ctrl := controller.NewUserService(&srv)

	return &Module{
		Controller: &ctrl,
	}
}
