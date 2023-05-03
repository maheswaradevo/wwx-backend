package authentications

import (
	"github.com/labstack/echo/v4"
	"github.com/maheswaradevo/wwx-backend/internal/model"
)

type AuthService interface {
	Login(ctx echo.Context, username, password string) (result *model.UserLoginResponse, err error)
}
