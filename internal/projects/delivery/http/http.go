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
		routeGroup.GET("/client/maintenance", projectDelivery.ViewClientMaintenanceProject)
		routeGroup.PATCH("/:projectId", projectDelivery.EditProject)
		routeGroup.POST("/search", projectDelivery.SearchProject)
		routeGroup.POST("/view", projectDelivery.ViewProject)
		routeGroup.GET("/client/view", projectDelivery.ViewClientProject)
		routeGroup.DELETE("/:projectId", projectDelivery.DeleteProject)
		routeGroup.GET("/view/edit/:projectId", projectDelivery.ViewEditProject)
	}
	return
}

// Project godoc
//
//		@Summary		Add project to the website
//		@Description	API Endpoint for adding project to the website
//		@Tags			Create Project
//		@Accept			json
//		@Produce		json
//		@Param			addProject	    body		 model.ProjectRequest	true	"createProject"
//	    @Param			Authorization	header	     string                 true     "Bearer Token"
//		@Success		200		                     {object}	model.Project
//		@Failure		400		                     {object}	common.APIError
//		@Failure		404		                     {object}	common.APIError
//		@Failure		500		                     {object}	common.APIError
//		@Router			/projects/ [post]
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

// Project godoc
//
//		@Summary		Edit a project
//		@Description	API Endpoint for editing a project from the website
//		@Tags			Edit Project
//		@Accept			json
//		@Produce		json
//		@Param			id	    path		 string	true	"project id"
//		@Param			editProject	    body		 model.EditProjectRequest	true	"Fill with project details"
//	    @Param			Authorization	header	     string                 true     "Bearer Token"
//		@Success		200		                     {object}	model.Project
//		@Failure		400		                     {object}	common.APIError
//		@Failure		404		                     {object}	common.APIError
//		@Failure		500		                     {object}	common.APIError
//		@Router			/projects/{id} [patch]
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

// SearchProject godoc
//
//		@Summary		Search for a project
//		@Description	API Endpoint for searching project in website
//		@Tags			Search Project
//		@Accept			json
//		@Produce		json
//		@Param			projectName	    query		 string	true	"Project that want to be searched"
//	    @Param			Authorization	header	     string                 true     "Bearer Token"
//		@Success		200		                     {array}	model.Project
//		@Failure		400		                     {object}	common.APIError
//		@Failure		404		                     {object}	common.APIError
//		@Failure		500		                     {object}	common.APIError
//		@Router			/projects/search [post]
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

// ViewProject godoc
//
//		@Summary		View Project that exists in the website
//		@Description	API Endpoint for view all of the project
//		@Tags			View Project
//		@Accept			json
//		@Produce		json
//	    @Param			Authorization	header	     string                 true     "Bearer Token"
//	    @Param			json	body	     model.ProjectViewRequest                 true     "Fill with project status"
//		@Success		200		                     {array}	model.Project
//		@Failure		400		                     {object}	common.APIError
//		@Failure		404		                     {object}	common.APIError
//		@Failure		500		                     {object}	common.APIError
//		@Router			/projects/view [post]
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

// ViewMaintenanceProject godoc
//
//		@Summary		View project that on maintenance status
//		@Description	API Endpoint for view the project that's on maintenance status
//		@Tags			View Project Maintenance
//		@Accept			json
//		@Produce		json
//	    @Param			Authorization	header	     string                 true     "Bearer Token"
//		@Success		200		                     {object}	model.Project
//		@Failure		400		                     {object}	common.APIError
//		@Failure		404		                     {object}	common.APIError
//		@Failure		500		                     {object}	common.APIError
//		@Router			/maintenance [get]
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

// ViewClientMaintenanceProject godoc
//
//		@Summary		View client project that on maintenance status
//		@Description	API Endpoint for view the client project that's on maintenance status
//		@Tags			View Client Project Maintenance
//		@Accept			json
//		@Produce		json
//	    @Param			Authorization	header	     string                 true     "Bearer Token"
//		@Success		200		                     {array}	model.Project
//		@Failure		400		                     {object}	common.APIError
//		@Failure		404		                     {object}	common.APIError
//		@Failure		500		                     {object}	common.APIError
//		@Router			/client/maintenance [get]
func (h ProjectHTTPDelivery) ViewClientMaintenanceProject(ctx echo.Context) error {
	user := ctx.Get("userData").(jwt.MapClaims)
	username := user["username"].(string)
	res, err := h.projectService.ViewClientMaintenanceProject(ctx, username)
	if err != nil {
		return h.InternalServerError(ctx, &common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
		})
	}
	return h.Ok(ctx, res)
}

// CreateMaintenanceProject godoc
//
//		@Summary		Add project that needs to be maintained to the website
//		@Description	API Endpoint for adding project to the website that need to maintained
//		@Tags			Create Maintenance Project
//		@Accept			json
//		@Produce		json
//		@Param			addProject	    body		 model.ProjectRequest	true	"createProject"
//	    @Param			Authorization	header	     string                 true     "Bearer Token"
//		@Success		200		                     {object}	model.Project
//		@Failure		400		                     {object}	common.APIError
//		@Failure		404		                     {object}	common.APIError
//		@Failure		500		                     {object}	common.APIError
//		@Router			/maintenance [post]
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

// DeleteProject godoc
//
//		@Summary		Delete a project
//		@Description	API Endpoint for deleting specified project by it's id
//		@Tags			Delete Project
//		@Accept			json
//		@Produce		json
//		@Param			id	    path		 string	true	"deleteProject"
//	    @Param			Authorization	header	     string                 true     "Bearer Token"
//		@Success		200		                     {object}	common.APIResponse
//		@Failure		400		                     {object}	common.APIError
//		@Failure		404		                     {object}	common.APIError
//		@Failure		500		                     {object}	common.APIError
//		@Router			/projects/{id} [delete]
func (h ProjectHTTPDelivery) DeleteProject(ctx echo.Context) error {
	projectId := ctx.Param("projectId")
	projectIdConv, _ := strconv.Atoi(projectId)

	err := h.projectService.DeleteProject(ctx, projectIdConv)
	if err != nil {
		h.logger.Sugar().Errorf("[DeleteProject] failed to delete project: %v", err)
		return h.InternalServerError(ctx, &common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
		})
	}
	return h.Ok(ctx, common.APIResponse{
		Code:    http.StatusOK,
		Message: "Deleted",
	})
}

// ViewClientProject godoc
//
//		@Summary		View client project
//		@Description	API Endpoint for view all of the client project based on their username
//		@Tags			View Client Project
//		@Accept			json
//		@Produce		json
//	    @Param			Authorization	header	     string                 true     "Bearer Token"
//		@Success		200		                     {array}	model.Project
//		@Failure		400		                     {object}	common.APIError
//		@Failure		404		                     {object}	common.APIError
//		@Failure		500		                     {object}	common.APIError
//		@Router			/projects/client/view [get]
func (h ProjectHTTPDelivery) ViewClientProject(ctx echo.Context) error {
	user := ctx.Get("userData").(jwt.MapClaims)
	username := user["username"].(string)

	res, err := h.projectService.ViewClientProject(ctx, username)
	if err != nil {
		h.logger.Sugar().Errorf("[ViewClientProject] failed to view client project: %v", err)
		return h.InternalServerError(ctx, &common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
		})
	}
	return h.Ok(ctx, res)
}

// ViewEditProject godoc
//
//		@Summary		View edited project in website
//		@Description	API Endpoint for view all of the data that used in the edit section on the website
//		@Tags			View Edit Project
//		@Accept			json
//		@Produce		json
//	    @Param			Authorization	header	     string                 true     "Bearer Token"
//		@Param			id	    path		 string	true	"viewEditProject"
//		@Success		200		                     {array}	model.Project
//		@Failure		400		                     {object}	common.APIError
//		@Failure		404		                     {object}	common.APIError
//		@Failure		500		                     {object}	common.APIError
//		@Router			/projects/view/edit/{id} [get]
func (h ProjectHTTPDelivery) ViewEditProject(ctx echo.Context) error {
	projectId := ctx.Param("projectId")
	projectIdConv, _ := strconv.Atoi(projectId)

	res, err := h.projectService.ViewEditProject(ctx, projectIdConv)
	if err != nil {
		h.logger.Sugar().Errorf("[ViewEditProject] failed to view client project: %v", err)
		return h.InternalServerError(ctx, &common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
		})
	}
	return h.Ok(ctx, res)
}
