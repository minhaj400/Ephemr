package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Minhajxdd/Ephemr/internal/auth/dto"
	"github.com/Minhajxdd/Ephemr/internal/auth/errors"
	authmodel "github.com/Minhajxdd/Ephemr/internal/auth/model"
	"github.com/Minhajxdd/Ephemr/internal/auth/repository"
	"github.com/Minhajxdd/Ephemr/internal/auth/utils"
	"github.com/Minhajxdd/Ephemr/internal/config"
	"github.com/Minhajxdd/Ephemr/internal/user/model"

	repositories "github.com/Minhajxdd/Ephemr/internal/user/repository"
	"github.com/Minhajxdd/Ephemr/pkg/crypto"
	"github.com/Minhajxdd/Ephemr/pkg/errs"
	"github.com/Minhajxdd/Ephemr/pkg/jwt"

	"github.com/google/uuid"
)

type AuthService interface {
	SignUp(user *dto.SignUpRequest) (*model.User, error)
}

type authService struct {
	userRepo       repositories.UserRepository
	emailUtils     utils.AuthEmailUtils
	emailTokenRepo repository.EmailTokenRepository
	hasher         crypto.PasswordHasher
	jwtManager     jwt.TokenManager
}

func NewAuthService(r repositories.UserRepository, e utils.AuthEmailUtils, er repository.EmailTokenRepository, h crypto.PasswordHasher, j jwt.TokenManager) AuthService {
	return &authService{userRepo: r, emailUtils: e, emailTokenRepo: er, hasher: h, jwtManager: j}
}

func (s *authService) SignUp(user *dto.SignUpRequest) (*model.User, error) {
	existingUser, err := s.userRepo.FindByEmail(user.Email)

	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errs.New(
			errors.UserDuplicateEmail,
			"user with email already exists",
			400,
			nil,
		)
	}

	hashedPassword, err := s.hasher.Hash(user.Password)
	if err != nil {
		return nil, errs.New(
			errors.InvalidPassword,
			"invalid password",
			400,
			nil,
		)
	}

	newUser := &model.User{
		Name:     user.FullName,
		Email:    user.Email,
		Password: hashedPassword,
	}

	err = s.userRepo.Create(newUser)

	if err != nil {
		return nil, errs.InternalError(err)
	}

	newUser.Password = ""

	emailTokenHash := uuid.NewString()

	emailToken := &authmodel.EmailToken{
		UserID:    newUser.ID,
		TokenHash: emailTokenHash,
		Kind:      authmodel.TokenKindVerify,
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}

	err = s.emailTokenRepo.Create(emailToken)

	if err != nil {
		return nil, errs.InternalError(err)
	}

	userId := strconv.FormatUint(uint64(newUser.ID), 10)

	magicLink := fmt.Sprintf("%s/verify-email/%s/%s", config.Cfg.HostName, userId, emailTokenHash)

	s.emailUtils.SentConfirmEmail(newUser.Email, magicLink)

	return newUser, nil
}
