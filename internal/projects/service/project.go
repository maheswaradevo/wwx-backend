package service

import (
	"github.com/labstack/echo/v4"
	"github.com/maheswaradevo/wwx-backend/internal/model"
	"github.com/maheswaradevo/wwx-backend/internal/projects"
	"github.com/maheswaradevo/wwx-backend/pkg/common/constants"
	"github.com/maheswaradevo/wwx-backend/pkg/common/helpers"
	"go.uber.org/zap"
)

type service struct {
	repo   projects.ProjectRepository
	logger *zap.Logger
}

func NewProjetService(repo projects.ProjectRepository, logger *zap.Logger) *service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) InsertProject(ctx echo.Context, dataRequest model.ProjectRequest) (*model.Project, error) {
	result, err := s.repo.InsertProject(helpers.Context(ctx), model.Project{
		ProjectName:  dataRequest.ProjectName,
		ClientName:   dataRequest.ClientName,
		Resource:     dataRequest.Resource,
		Deadline:     dataRequest.Deadline,
		Status:       dataRequest.Status,
		ProposalLink: dataRequest.ProposalLink,
		Assign:       dataRequest.Assign,
	})
	if err != nil {
		s.logger.Sugar().Errorf("[InsertProject] failed to insert project: %v", zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (s *service) EditProject(ctx echo.Context, dataRequest model.EditProjectRequest, projectId int, role string) (*model.Project, error) {
	exist, err := s.repo.CheckProject(helpers.Context(ctx), projectId)
	if err != nil {
		s.logger.Sugar().Errorf("[EditProject] failed to check product with, err: %v", zap.Error(err))
		return nil, err
	}
	if !exist {
		err = constants.ErrDataNotFound
		s.logger.Sugar().Errorf("[EditProject] project not found, err: %v", zap.Error(err))
		return nil, err
	}
	if role == constants.RoleAdmin {
		err := s.repo.UpdateProjectAdmin(helpers.Context(ctx), dataRequest, projectId)
		if err != nil {
			s.logger.Sugar().Errorf("[EditProject] failed to update project as admin: %v", zap.Error(err))
			return nil, err
		}
		var resp = &model.Project{
			ProjectID:    projectId,
			ProjectName:  dataRequest.ProjectName,
			ClientName:   dataRequest.ClientName,
			Deadline:     dataRequest.Deadline,
			Status:       dataRequest.Status,
			Budget:       dataRequest.Budget,
			ProposalLink: dataRequest.ProposalLink,
			Resource:     dataRequest.Resource,
			Assign:       dataRequest.Assign,
		}
		return resp, nil
	} else if role == constants.RoleClient {
		err := s.repo.UpdateProjectClient(helpers.Context(ctx), dataRequest, projectId)
		if err != nil {
			s.logger.Sugar().Errorf("[EditProject] failed to update project as client: %v", zap.Error(err))
			return nil, err
		}
		var resp = &model.Project{
			ProjectID: projectId,
			Budget:    dataRequest.Budget,
			Resource:  dataRequest.Resource,
		}
		return resp, nil
	}
	return nil, nil
}
