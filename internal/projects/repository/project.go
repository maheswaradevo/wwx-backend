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
	res, err := stmt.ExecContext(ctx, data.ProjectName, data.ClientName, data.Resource, data.Deadline, data.Status, data.ProposalLink, data.Assign)
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

func (p projectRepository) CheckProject(ctx context.Context, projectId int) (bool, error) {
	query := constants.CheckProject
	rows, err := p.db.QueryContext(ctx, query, projectId)
	if err != nil {
		p.logger.Sugar().Errorf("[CheckProject] failed to query to the database: %v", zap.Error(err))
		return false, err
	}
	for rows.Next() {
		return true, nil
	}
	return false, nil
}

func (p projectRepository) UpdateProjectAdmin(ctx context.Context, data model.EditProjectRequest, projectId int) error {
	query := constants.UpdateProjectAdmin
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		p.logger.Sugar().Errorf("[UpdateProjectAdmin] failed to prepare the statement: %v", zap.Error(err))
		return err
	}
	_, err = stmt.ExecContext(ctx, data.ProjectName, data.ClientName, data.Deadline, data.Status, data.Budget, data.ProposalLink, data.Assign, data.Resource, projectId)
	if err != nil {
		p.logger.Sugar().Errorf("[UpdateProjectAdmin] failed to insert data to the database: %v", zap.Error(err))
		return err
	}
	return nil
}

func (p projectRepository) UpdateProjectClient(ctx context.Context, data model.EditProjectRequest, projectId int) error {
	query := constants.UpdateProjectClient
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		p.logger.Sugar().Errorf("[UpdateProjectClient] failed to prepare the statement: %v", zap.Error(err))
		return err
	}
	_, err = stmt.ExecContext(ctx, data.Budget, data.Resource, projectId)
	if err != nil {
		p.logger.Sugar().Errorf("[UpdateProjectClient] failed to insert data to the database: %v", zap.Error(err))
		return err
	}
	return nil
}
