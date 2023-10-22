package projects

import (
	"context"

	"github.com/maheswaradevo/wwx-backend/internal/model"
)

type ProjectRepository interface {
	InsertProject(ctx context.Context, data model.Project, userId int) (*model.Project, error)
	InsertMaintenanceProject(ctx context.Context, data model.Project, userId int) (*model.Project, error)
	ViewMaintenanceProject(ctx context.Context) (res []*model.Project, err error)
	CheckProject(ctx context.Context, projectId int) (bool, error)
	UpdateProjectAdmin(ctx context.Context, data model.EditProjectRequest, projectId int) error
	UpdateProjectClient(ctx context.Context, data model.EditProjectRequest, projectId int) error
	SearchProject(ctx context.Context, projectName string) (projects []*model.Project, err error)
	ViewAdminProject(ctx context.Context, userId int, status string) (res []*model.Project, err error)
	DeleteProject(ctx context.Context, projectId int) error
	ViewClientMaintenanceProject(ctx context.Context) (res []*model.Project, err error)
	ViewClientProject(ctx context.Context, username string) (res []*model.Project, err error)
	ViewEditProject(ctx context.Context, projectId int) (res []*model.Project, err error)
}
