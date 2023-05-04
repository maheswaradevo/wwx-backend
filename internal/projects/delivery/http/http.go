package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maheswaradevo/wwx-backend/internal/model"
	"github.com/maheswaradevo/wwx-backend/internal/projects"
	"github.com/maheswaradevo/wwx-backend/pkg/common"
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
	}
	return
}

func (h ProjectHTTPDelivery) CreateProject(ctx echo.Context) error {
	var req model.ProjectRequest

	if err := ctx.Bind(&req); err != nil {
		return h.WrapBadRequest(ctx, &common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Errors:  "error binding",
		})
	}

	result, err := h.projectService.InsertProject(ctx, req)
	if err != nil {
		h.logger.Sugar().Errorf("[createProject] failed to create project, err: %v", err)
		return h.InternalServerError(ctx, &common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}
	return h.Ok(ctx, result)
}
