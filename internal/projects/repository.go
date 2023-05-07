package projects

import (
	"context"

	"github.com/maheswaradevo/wwx-backend/internal/model"
)

type ProjectRepository interface {
	InsertProject(ctx context.Context, data model.Project, userId int) (*model.Project, error)
	InsertMaintenanceProject(ctx context.Context, data model.Project, userId int) (*model.Project, error)
	CheckProject(ctx context.Context, projectId int) (bool, error)
	UpdateProjectAdmin(ctx context.Context, data model.EditProjectRequest, projectId int) error
	UpdateProjectClient(ctx context.Context, data model.EditProjectRequest, projectId int) error
	SearchProject(ctx context.Context, projectName string) (projects []*model.Project, err error)
	ViewProject(ctx context.Context, userId int) (res []*model.Project, err error)
}
