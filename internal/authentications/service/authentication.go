package service

import (
	"github.com/labstack/echo/v4"
	"github.com/maheswaradevo/wwx-backend/internal/authentications"
	"github.com/maheswaradevo/wwx-backend/internal/model"
	"github.com/maheswaradevo/wwx-backend/pkg/common/constants"
	"github.com/maheswaradevo/wwx-backend/pkg/common/helpers"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo   authentications.AuthRepository
	logger *zap.Logger
}

func NewAuthService(repo authentications.AuthRepository, logger *zap.Logger) *service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (auth *service) Login(ctx echo.Context, username, password string) (result *model.UserLoginResponse, err error) {
	userCred, errFindByUsername := auth.repo.FindByUsername(helpers.Context(ctx), username)
	if err != nil {
		auth.logger.Sugar().Errorf("[Login] failed to fetch data by email, err: %v", zap.Error(errFindByUsername))
		return nil, errFindByUsername
	}
	errMismatchPassword := bcrypt.CompareHashAndPassword([]byte(userCred.Password), []byte(password))
	if errMismatchPassword != nil {
		errMismatchPassword = constants.ErrMismatchedHashAndPassword
		auth.logger.Sugar().Errorf("[Login] wrong password")
		return nil, errMismatchPassword
	}
	token, errCreateAccessToken := helpers.CreateAccessToken(userCred)
	if errCreateAccessToken != nil {
		auth.logger.Sugar().Errorf("[Login] failed to create access token, err: %v", zap.Error(errCreateAccessToken))
		return nil, errCreateAccessToken
	}

	var res = model.UserLoginResponse{
		ID:          userCred.ID,
		Username:    userCred.Username,
		Role:        userCred.Role,
		AccessToken: token,
	}
	return &res, err
}
