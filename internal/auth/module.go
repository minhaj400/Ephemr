package auth

import (
	"github.com/Minhajxdd/Ephemr/internal/auth/controller"
	"github.com/Minhajxdd/Ephemr/internal/auth/repository"
	"github.com/Minhajxdd/Ephemr/internal/auth/service"
	"github.com/Minhajxdd/Ephemr/internal/auth/utils"
	repositories "github.com/Minhajxdd/Ephemr/internal/user/repository"
	"github.com/Minhajxdd/Ephemr/pkg/crypto"
	"github.com/Minhajxdd/Ephemr/pkg/jwt"
	"gorm.io/gorm"
)

type Module struct {
	AuthController controller.AuthController
}

func NewModule(r repositories.UserRepository, db *gorm.DB, h crypto.PasswordHasher, j jwt.TokenManager) *Module {
	emailUtils := utils.NewAuthEmailUtils()
	emailTokenRepo := repository.NewEmailTokenRepository(db)
	refreshTokenRepo := repository.NewRefreshTokenRepo(db)
	srv := service.NewAuthService(r, emailUtils, emailTokenRepo, h, j, refreshTokenRepo)
	ctrl := controller.NewAuthController(srv)

	return &Module{
		AuthController: ctrl,
	}
}
