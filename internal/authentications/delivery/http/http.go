package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maheswaradevo/wwx-backend/internal/authentications"
	"github.com/maheswaradevo/wwx-backend/internal/model"
	"github.com/maheswaradevo/wwx-backend/pkg/common"
	"github.com/maheswaradevo/wwx-backend/pkg/common/constants"
	"go.uber.org/zap"
)

type AuthenticationHTTPDelivery struct {
	common.Controller
	authService  authentications.AuthService
	routeGroupV1 *echo.Group
	logger       *zap.Logger
}

func AuthenticationNewDelivery(authService authentications.AuthService, routeGroupV1 *echo.Group, logger *zap.Logger) (routeGroup *echo.Group) {
	authenticationDelivery := AuthenticationHTTPDelivery{
		authService:  authService,
		routeGroupV1: routeGroupV1,
		logger:       logger,
	}
	routeGroup = authenticationDelivery.routeGroupV1.Group("/auth")
	{
		routeGroup.POST("/login", authenticationDelivery.Login)
	}
	return
}

func (h AuthenticationHTTPDelivery) Login(ctx echo.Context) error {
	var req model.UserLoginRequest

	if err := ctx.Bind(&req); err != nil {
		return h.WrapBadRequest(ctx, &common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Errors:  "error binding",
		})
	}

	// if err := ctx.Validate(&req); err != nil {
	// 	return h.WrapBadRequest(ctx, &common.APIResponse{
	// 		Code:    http.StatusBadRequest,
	// 		Message: http.StatusText(http.StatusBadRequest),
	// 		Errors:  "error validating",
	// 	})
	// }

	result, err := h.authService.Login(ctx, req.Username, req.Password)
	if err != nil {
		h.logger.Sugar().Errorf("[login] failed to login, err: %v", err)
		return h.WrapBadRequest(ctx, &common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Masukkan password yang sesuai",
		})
	}
	var resp = common.APIResponse{
		Code:    http.StatusOK,
		Message: constants.LoginSuccess,
		Data:    result,
	}
	return h.Ok(ctx, resp)
}
