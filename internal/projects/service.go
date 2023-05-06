package projects

import (
	"github.com/labstack/echo/v4"
	"github.com/maheswaradevo/wwx-backend/internal/model"
)

type ProjectService interface {
	InsertProject(ctx echo.Context, dataRequest model.ProjectRequest, role string, userId int) (res *model.Project, err error)
	EditProject(ctx echo.Context, dataRequest model.EditProjectRequest, projectId int, role string) (*model.Project, error)
	SearchProject(ctx echo.Context, projectName string) (projects []*model.Project, err error)
	ViewProject(ctx echo.Context, userRole int) (projects []*model.Project, err error)
}
