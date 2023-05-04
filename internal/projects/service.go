package projects

import (
	"github.com/labstack/echo/v4"
	"github.com/maheswaradevo/wwx-backend/internal/model"
)

type ProjectService interface {
	InsertProject(ctx echo.Context, dataRequest model.ProjectRequest) (*model.Project, error)
	EditProject(ctx echo.Context, dataRequest model.EditProjectRequest, projectId int, role string) (*model.Project, error)
}
