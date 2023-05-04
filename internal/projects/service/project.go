package service

import (
	"github.com/labstack/echo/v4"
	"github.com/maheswaradevo/wwx-backend/internal/model"
	"github.com/maheswaradevo/wwx-backend/internal/projects"
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
