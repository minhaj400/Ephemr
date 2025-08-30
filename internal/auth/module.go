package auth

import (
	"github.com/Minhajxdd/Ephemr/internal/auth/controller"
	"github.com/Minhajxdd/Ephemr/internal/auth/service"
	repositories "github.com/Minhajxdd/Ephemr/internal/user/repository"
	"github.com/Minhajxdd/Ephemr/pkg/crypto"
	"github.com/Minhajxdd/Ephemr/pkg/jwt"
)

type Module struct {
	AuthController *controller.AuthController
}

func NewModule(r *repositories.UserRepository, h crypto.PasswordHasher, j jwt.TokenManager) *Module {
	srv := service.NewAuthService(*r, h, j)
	ctrl := controller.NewAuthController(srv)

	return &Module{
		AuthController: &ctrl,
	}
}
