package service

import (
	"fmt"
	"net/http"
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

// AuthService defines the methods for handling the business logic of user authentication.
type AuthService interface {
	// SignUp handles the business logic for user registration, including creating a user record & sending confirmation mail.
	SignUp(user *dto.SignUpRequest) (*model.User, error)

	// ConfirmEmail handles the business logic for confirming a user's email address.
	ConfirmEmail(params *dto.ConfirmEmailRequest, device, ipAddress string) (string, string, error)

	// RefreshToken handles the business logic for refreshing access tokens.
	RefreshToken(token, device, ipAddress string) (string, string, error)

	// Login handles the business logic for user login, including credential validation.
	Login(user *dto.LoginRequest, device, ipAddress string) (string, string, error)
}

type authService struct {
	userRepo         repositories.UserRepository
	emailUtils       utils.AuthEmailUtils
	emailTokenRepo   repository.EmailTokenRepository
	refreshTokenRepo repository.RefreshTokenRepository
	hasher           crypto.PasswordHasher
	jwtManager       jwt.TokenManager
}

func NewAuthService(
	r repositories.UserRepository,
	e utils.AuthEmailUtils,
	er repository.EmailTokenRepository,
	h crypto.PasswordHasher,
	j jwt.TokenManager,
	rk repository.RefreshTokenRepository) AuthService {

	return &authService{userRepo: r, emailUtils: e, emailTokenRepo: er, hasher: h, jwtManager: j, refreshTokenRepo: rk}
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

	magicLink := fmt.Sprintf("%s/confirm-email/%s/%s", config.Cfg.HostName, userId, emailTokenHash)

	s.emailUtils.SentConfirmEmail(newUser.Email, magicLink)

	return newUser, nil
}

func (s *authService) Login(u *dto.LoginRequest, device, ipAddress string) (string, string, error) {
	user, err := s.userRepo.FindByEmail(u.Email)

	if err != nil {
		return "", "", errs.InternalError(err)
	}

	if user == nil || !user.IsVerified {
		return "", "", errs.New(
			errors.InvalidPasswordOrUser,
			"Invalid Password Or User",
			http.StatusBadRequest,
			nil,
		)
	}

	if isValidPassword := s.hasher.Compare(u.Password, user.Password); !isValidPassword {
		return "", "", errs.New(
			errors.InvalidPasswordOrUser,
			"Invalid Password Or User",
			http.StatusBadRequest,
			nil,
		)
	}

	userId := strconv.FormatUint(uint64(user.ID), 10)
	var claims = jwt.Claims{
		UserID: userId,
		Role:   "user",
	}

	jwtToken, err := s.jwtManager.Generate(claims)

	if err != nil {
		return "", "", errs.InternalError(err)
	}

	refreshTokenHash := uuid.NewString()

	refreshToken := &authmodel.RefreshTokens{
		UserID:    user.ID,
		TokenHash: refreshTokenHash,
		Device:    device,
		IpAddress: ipAddress,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	err = s.refreshTokenRepo.Create(refreshToken)

	if err != nil {
		return "", "", errs.InternalError(err)
	}

	return jwtToken, refreshTokenHash, nil
}

func (s *authService) ConfirmEmail(params *dto.ConfirmEmailRequest, device, ipAddress string) (string, string, error) {
	token, err := s.emailTokenRepo.IsValid(params.UserId, params.Token, authmodel.TokenKindVerify)

	if err != nil {
		return "", "", errs.InternalError(err)
	}

	if token == nil {
		return "", "", errs.New(
			errors.InvalidToken,
			"invalid token",
			400,
			nil,
		)
	}

	currentTime := time.Now()
	isExpired := token.ExpiresAt.Before(currentTime)

	if isExpired {
		if err := s.emailTokenRepo.DeleteById(token.ID); err != nil {
			return "", "", errs.InternalError(err)
		}

		return "", "", errs.New(
			errors.TokenExpired,
			"token expired",
			400,
			nil,
		)
	}

	if err := s.emailTokenRepo.DeleteById(token.ID); err != nil {
		return "", "", errs.InternalError(err)
	}

	updatedUser, err := s.userRepo.SetVerifyStatus(params.UserId, true)

	if err != nil {
		return "", "", errs.InternalError(err)
	}

	userId := strconv.FormatUint(uint64(updatedUser.ID), 10)
	var claims = jwt.Claims{
		UserID: userId,
		Role:   "user",
	}

	jwtToken, err := s.jwtManager.Generate(claims)

	if err != nil {
		return "", "", errs.InternalError(err)
	}

	refreshTokenHash := uuid.NewString()

	refreshToken := &authmodel.RefreshTokens{
		UserID:    updatedUser.ID,
		TokenHash: refreshTokenHash,
		Device:    device,
		IpAddress: ipAddress,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	err = s.refreshTokenRepo.Create(refreshToken)

	if err != nil {
		return "", "", errs.InternalError(err)
	}

	return jwtToken, refreshTokenHash, nil
}

func (s *authService) RefreshToken(token, device, ipAddress string) (string, string, error) {
	refreshToken, err := s.refreshTokenRepo.FindWithTokenDeviceIp(token, device, ipAddress)

	if err != nil {
		return "", "", err
	}

	if refreshToken == nil {
		return "", "", errs.New(
			errors.TokenExpired,
			"token expired",
			http.StatusUnauthorized,
			nil,
		)
	}

	currentTime := time.Now()
	isExpired := refreshToken.ExpiresAt.Before(currentTime)

	if isExpired {
		if err := s.refreshTokenRepo.DeleteById(refreshToken.ID); err != nil {
			return "", "", errs.InternalError(err)
		}

		return "", "", errs.New(
			errors.TokenExpired,
			"token expired",
			http.StatusUnauthorized,
			nil,
		)
	}

	user, err := s.userRepo.GetByID(refreshToken.UserID)

	if err != nil {
		return "", "", err
	}

	if user == nil {
		return "", "", errs.BadRequest("Invalid Token", nil)
	}

	userId := strconv.FormatUint(uint64(user.ID), 10)
	var claims = jwt.Claims{
		UserID: userId,
		Role:   "user",
	}

	jwtToken, err := s.jwtManager.Generate(claims)

	if err != nil {
		return "", "", errs.InternalError(err)
	}

	refreshTokenHash := uuid.NewString()

	refreshToken.TokenHash = refreshTokenHash
	refreshToken.ExpiresAt = time.Now().Add(7 * 24 * time.Hour)

	err = s.refreshTokenRepo.Update(refreshToken)

	if err != nil {
		return "", "", errs.InternalError(err)
	}

	return jwtToken, refreshTokenHash, nil
}
