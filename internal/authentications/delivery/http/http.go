package http

import (
	"errors"
	"fmt"
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
			Errors:  constants.BindingRequestError,
		})
	}

	result, err := h.authService.Login(ctx, req.Username, req.Password)
	if err != nil {
		h.logger.Sugar().Errorf("[login] failed to login, err: %v", err)
		switch {
		case errors.Is(constants.ErrMismatchedHashAndPassword, err):
			respError := common.APIError{
				Code:    fmt.Sprint(http.StatusBadRequest),
				Field:   "Credentials",
				Message: constants.PasswordMismatch,
			}
			apiError := respError.SetInternal(err)
			return ctx.JSON(http.StatusBadRequest, apiError)

		case errors.Is(constants.ErrNoUsernameExist, err):
			respError := common.APIError{
				Code:    fmt.Sprint(http.StatusBadRequest),
				Field:   "Username",
				Message: constants.NoUsernameExists,
			}
			apiError := respError.SetInternal(err)
			return ctx.JSON(http.StatusBadRequest, apiError)
		default:
			return h.InternalServerError(ctx, &common.APIResponse{
				Code:    http.StatusInternalServerError,
				Message: constants.InternalServerError,
			})
		}
	}
	return h.Ok(ctx, result)
}
