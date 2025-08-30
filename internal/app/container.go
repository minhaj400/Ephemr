package app

import (
	"time"

	"github.com/Minhajxdd/Ephemr/config"
	"github.com/Minhajxdd/Ephemr/internal/auth"
	"github.com/Minhajxdd/Ephemr/internal/user"
	"github.com/Minhajxdd/Ephemr/pkg/crypto"
	"github.com/Minhajxdd/Ephemr/pkg/jwt"
	"gorm.io/gorm"
)

type Container struct {
	DB         *gorm.DB
	Hasher     crypto.PasswordHasher
	JWTManager jwt.TokenManager

	UserModule *user.Module
	AuthModule *auth.Module
}

func NewContainer() *Container {
	gormDB := config.DB

	hasher := crypto.NewBcryptHasher(14)
	jwtMgr := jwt.NewJWTManager(config.Cfg.JwtSecret, time.Minute*30)

	userMod := user.NewModule(gormDB)
	// authMod := auth.NewModule(gormDB) // auth depends on user repo

	return &Container{
		DB:         gormDB,
		Hasher:     hasher,
		JWTManager: jwtMgr,
		UserModule: userMod,
		// AuthModule: authMod,
	}
}
