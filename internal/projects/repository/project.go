package repository

import (
	"context"
	"database/sql"

	"github.com/maheswaradevo/wwx-backend/internal/model"
	"github.com/maheswaradevo/wwx-backend/pkg/common/constants"
	"go.uber.org/zap"
)

type projectRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewProjectRepository(db *sql.DB, logger *zap.Logger) *projectRepository {
	return &projectRepository{
		db:     db,
		logger: logger,
	}
}

func (p projectRepository) InsertProject(ctx context.Context, data model.Project) (*model.Project, error) {
	query := constants.InsertProject
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		p.logger.Sugar().Errorf("[InsertProject] failed to prepare statement: %v", zap.Error(err))
		return nil, err
	}
	res, err := stmt.ExecContext(ctx, data.ProjectName, data.ClientName, data.Deadline, data.Status, data.ProposalLink, data.Assign)
	if err != nil {
		p.logger.Sugar().Errorf("[InsertProject] failed to insert user to the database: %v", zap.Error(err))
	}
	id, _ := res.LastInsertId()
	prj := model.Project{
		ProjectID:    int(id),
		ProjectName:  data.ProjectName,
		ClientName:   data.ClientName,
		Deadline:     data.Deadline,
		Status:       data.Status,
		ProposalLink: data.ProposalLink,
		Assign:       data.Assign,
	}

	return &prj, nil
}
