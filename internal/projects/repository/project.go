package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

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

func (p projectRepository) InsertProject(ctx context.Context, data model.Project, userId int) (*model.Project, error) {
	query := constants.InsertProject
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		p.logger.Sugar().Errorf("[InsertProject] failed to prepare statement: %v", zap.Error(err))
		return nil, err
	}
	res, err := stmt.ExecContext(ctx, data.UserId, data.ProjectName, data.ClientName, data.Resource, data.Deadline, data.Status, data.ProposalLink, data.Assign, data.Budget)
	if err != nil {
		p.logger.Sugar().Errorf("[InsertProject] failed to insert user to the database: %v", zap.Error(err))
	}
	id, _ := res.LastInsertId()
	prj := model.Project{
		UserId:       userId,
		ProjectID:    int(id),
		ProjectName:  data.ProjectName,
		ClientName:   data.ClientName,
		Deadline:     data.Deadline,
		Status:       data.Status,
		ProposalLink: data.ProposalLink,
		Assign:       data.Assign,
		Budget:       data.Budget,
		CreatedAt:    time.Now(),
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

func (p projectRepository) SearchProject(ctx context.Context, projectName string) (projects []*model.Project, err error) {
	query := constants.SearchProject
	lowerProjectName := strings.ToLower(projectName)
	queryRes := fmt.Sprintf(query, lowerProjectName)

	stmt, err := p.db.PrepareContext(ctx, queryRes)
	if err != nil {
		p.logger.Sugar().Errorf("[SearchProject] failed to prepare the statement: %v", zap.Error(err))
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		p.logger.Sugar().Errorf("[SearchProject] failed to query to the database: %v", zap.Error(err))
		return nil, err
	}
	var project model.Project
	for rows.Next() {
		err := rows.Scan(
			&project.ProjectID,
			&project.ProjectName,
			&project.ClientName,
			&project.Deadline,
			&project.Status,
			&project.Budget,
			&project.ProposalLink,
			&project.Assign,
			&project.UserId,
			&project.Resource,
			&project.CreatedAt,
		)
		if err != nil {
			p.logger.Sugar().Errorf("[SearhProject] failed to scan data from database: %v", zap.Error(err))
			return nil, err
		}
		projects = append(projects, &project)
	}
	return projects, nil
}

func (p projectRepository) ViewProject(ctx context.Context, userId int, status string) (res []*model.Project, err error) {
	var query string
	if userId == 1 {
		query = constants.ViewProjectAdmin
		rowsAdmin, err := p.db.QueryContext(ctx, query, status)
		if err != nil {
			p.logger.Sugar().Errorf("[ViewProjectAdmin] failed to query to the database: %v", zap.Error(err))
			return nil, err
		}

		for rowsAdmin.Next() {
			prj := model.Project{}
			err := rowsAdmin.Scan(
				&prj.ProjectID,
				&prj.ProjectName,
				&prj.ClientName,
				&prj.Deadline,
				&prj.Status,
				&prj.Budget,
				&prj.ProposalLink,
				&prj.Assign,
				&prj.UserId,
				&prj.Resource,
				&prj.CreatedAt,
				&prj.Maintenance,
			)
			if err != nil {
				p.logger.Sugar().Errorf("[ViewProjectAdmin] failed to scan the data: %v", zap.Error(err))
				return nil, err
			}
			res = append(res, &prj)
		}
	} else {
		query = constants.ViewProject
		rowsClient, err := p.db.QueryContext(ctx, query, userId)
		if err != nil {
			p.logger.Sugar().Errorf("[ViewProjectAdmin] failed to query to the database: %v", zap.Error(err))
			return nil, err
		}

		for rowsClient.Next() {
			prj := model.Project{}
			err := rowsClient.Scan(
				&prj.ProjectID,
				&prj.ProjectName,
				&prj.ClientName,
				&prj.Deadline,
				&prj.Status,
				&prj.Budget,
				&prj.ProposalLink,
				&prj.Assign,
				&prj.Resource,
				&prj.UserId,
			)
			if err != nil {
				p.logger.Sugar().Errorf("[ViewProjectAdmin] failed to scan the data: %v", zap.Error(err))
				return nil, err
			}
			res = append(res, &prj)
		}
	}

	return res, nil
}

func (p projectRepository) InsertMaintenanceProject(ctx context.Context, data model.Project, userId int) (*model.Project, error) {
	query := constants.InsertMaintenanceProject
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		p.logger.Sugar().Errorf("[InsertProject] failed to prepare statement: %v", zap.Error(err))
		return nil, err
	}
	res, err := stmt.ExecContext(ctx, data.UserId, data.ProjectName, data.ClientName, data.Resource, data.Deadline, data.Status, data.ProposalLink, data.Assign, data.Budget, data.Maintenance)
	if err != nil {
		p.logger.Sugar().Errorf("[InsertProject] failed to insert user to the database: %v", zap.Error(err))
	}
	id, _ := res.LastInsertId()
	prj := model.Project{
		UserId:       userId,
		ProjectID:    int(id),
		ProjectName:  data.ProjectName,
		ClientName:   data.ClientName,
		Deadline:     data.Deadline,
		Status:       data.Status,
		ProposalLink: data.ProposalLink,
		Assign:       data.Assign,
		Budget:       data.Budget,
		Maintenance:  1,
		CreatedAt:    time.Now(),
	}

	return &prj, nil
}

func (p projectRepository) ViewMaintenanceProject(ctx context.Context) (res []*model.Project, err error) {
	query := constants.ViewMaintenanceProject
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		p.logger.Sugar().Errorf("[ViewMaintenanceProject] failed to query to the database: %v", zap.Error(err))
		return nil, err
	}

	for rows.Next() {
		prj := model.Project{}
		err := rows.Scan(
			&prj.ProjectID,
			&prj.ProjectName,
			&prj.ClientName,
			&prj.Deadline,
			&prj.Status,
			&prj.Budget,
			&prj.ProposalLink,
			&prj.Assign,
			&prj.UserId,
			&prj.Resource,
			&prj.CreatedAt,
			&prj.Maintenance,
		)
		if err != nil {
			p.logger.Sugar().Errorf("[ViewProjectAdmin] failed to scan the data: %v", zap.Error(err))
			return nil, err
		}
		res = append(res, &prj)
	}
	return res, nil
}

func (p projectRepository) DeleteProject(ctx context.Context, projectId int) error {
	query := constants.DeleteProject

	_, err := p.db.ExecContext(ctx, query, projectId)
	if err != nil {
		p.logger.Sugar().Errorf("[DeleteProject] failed to delete the data: %v", zap.Error(err))
		return err
	}
	return nil
}
