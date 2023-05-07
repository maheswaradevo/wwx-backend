package http

import (
	"errors"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/maheswaradevo/wwx-backend/internal/model"
	"github.com/maheswaradevo/wwx-backend/internal/projects"
	"github.com/maheswaradevo/wwx-backend/pkg/common"
	"github.com/maheswaradevo/wwx-backend/pkg/common/constants"
	"go.uber.org/zap"
)

type ProjectHTTPDelivery struct {
	common.Controller
	projectService projects.ProjectService
	routeGroupV1   *echo.Group
	logger         *zap.Logger
}

func ProjectNewDelivery(projectService projects.ProjectService, routeGroupV1 *echo.Group, logger *zap.Logger) (routeGroup *echo.Group) {
	projectDelivery := ProjectHTTPDelivery{
		projectService: projectService,
		routeGroupV1:   routeGroupV1,
		logger:         logger,
	}
	routeGroup = projectDelivery.routeGroupV1.Group("/projects")
	{
		routeGroup.POST("/", projectDelivery.CreateProject)
		routeGroup.POST("/maintenance", projectDelivery.CreateMaintenanceProject)
		routeGroup.GET("/maintenance", projectDelivery.ViewMaintenanceProject)
		routeGroup.PUT("/:projectId", projectDelivery.EditProject)
		routeGroup.POST("/search", projectDelivery.SearchProject)
		routeGroup.POST("/view", projectDelivery.ViewProject)
	}
	return
}

func (h ProjectHTTPDelivery) CreateProject(ctx echo.Context) error {
	var req model.ProjectRequest

	user := ctx.Get("userData").(jwt.MapClaims)
	role := user["userRole"].(string)
	userId := int(user["userId"].(float64))

	if role != constants.RoleAdmin {
		return h.Unauthorized(ctx, errors.New("Unauthorized"))
	}
	if err := ctx.Bind(&req); err != nil {
		return h.WrapBadRequest(ctx, &common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Errors:  constants.BindingRequestError,
		})
	}
	result, err := h.projectService.InsertProject(ctx, req, role, userId)
	if err != nil {
		h.logger.Sugar().Errorf("[createProject] failed to create project, err: %v", err)
		return h.InternalServerError(ctx, &common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
		})
	}
	return h.Ok(ctx, result)
}

func (h ProjectHTTPDelivery) EditProject(ctx echo.Context) error {
	var req model.EditProjectRequest

	if err := ctx.Bind(&req); err != nil {
		return h.WrapBadRequest(ctx, &common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Errors:  constants.BindingRequestError,
		})
	}

	user := ctx.Get("userData").(jwt.MapClaims)
	role := user["userRole"].(string)

	projectId := ctx.Param("projectId")
	projectIdConv, _ := strconv.Atoi(projectId)

	result, err := h.projectService.EditProject(ctx, req, projectIdConv, role)
	if err != nil {
		switch {
		case errors.Is(constants.ErrDataNotFound, err):
			return h.DataNotFound(ctx, &common.APIResponse{
				Code:    http.StatusNotFound,
				Message: constants.ProjectNotFound,
			})
		default:
			return h.InternalServerError(ctx, &common.APIResponse{
				Code:    http.StatusInternalServerError,
				Message: constants.InternalServerError,
			})
		}
	}
	return h.Ok(ctx, result)
}

func (h ProjectHTTPDelivery) SearchProject(ctx echo.Context) error {
	queryVar := ctx.QueryParams()
	projectName := queryVar.Get("projectName")
	result, err := h.projectService.SearchProject(ctx, projectName)
	if err != nil {
		return h.InternalServerError(ctx, &common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
		})
	}
	return h.Ok(ctx, result)
}

func (h ProjectHTTPDelivery) ViewProject(ctx echo.Context) error {
	var req model.ProjectViewRequest

	if err := ctx.Bind(&req); err != nil {
		return h.WrapBadRequest(ctx, &common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Errors:  constants.BindingRequestError,
		})
	}

	user := ctx.Get("userData").(jwt.MapClaims)
	userId := int(user["userId"].(float64))

	res, err := h.projectService.ViewProject(ctx, userId, req.Status)
	if err != nil {
		return h.InternalServerError(ctx, &common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
		})
	}
	return h.Ok(ctx, res)
}

func (h ProjectHTTPDelivery) ViewMaintenanceProject(ctx echo.Context) error {
	res, err := h.projectService.ViewMaintenanceProject(ctx)
	if err != nil {
		return h.InternalServerError(ctx, &common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
		})
	}
	return h.Ok(ctx, res)
}

func (h ProjectHTTPDelivery) CreateMaintenanceProject(ctx echo.Context) error {
	var req model.ProjectRequest

	user := ctx.Get("userData").(jwt.MapClaims)
	userId := int(user["userId"].(float64))

	if err := ctx.Bind(&req); err != nil {
		return h.WrapBadRequest(ctx, &common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Errors:  constants.BindingRequestError,
		})
	}
	result, err := h.projectService.InsertMaintenanceProject(ctx, req, userId)
	if err != nil {
		h.logger.Sugar().Errorf("[createProject] failed to create project, err: %v", err)
		return h.InternalServerError(ctx, &common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
		})
	}
	return h.Ok(ctx, result)
}
