package projects

import (
	"context"

	"github.com/maheswaradevo/wwx-backend/internal/model"
)

type ProjectRepository interface {
	InsertProject(ctx context.Context, data model.Project) (*model.Project, error)
}
