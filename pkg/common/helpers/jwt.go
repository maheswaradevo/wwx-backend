package helpers

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/maheswaradevo/wwx-backend/internal/model"
	"github.com/maheswaradevo/wwx-backend/pkg/config"
	"go.uber.org/zap"
)

func CreateAccessToken(user *model.User) (string, error) {
	cfg := config.GetConfig()

	claims := jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 8).Unix(),
		"userId":     user.ID,
		"userRole":   user.Role,
		"username":   user.Username,
	}
	token := jwt.NewWithClaims(cfg.JWTSigningMethod, claims)
	signedToken, errSignedString := token.SignedString([]byte(cfg.ApiSecretKey))
	if errSignedString != nil {
		zap.S().Errorf("failed to create new token, err: %v", zap.Error(errSignedString))
		return "", errSignedString
	}
	return signedToken, nil
}
