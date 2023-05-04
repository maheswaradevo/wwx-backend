package pkg

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	authHTTPDelivery "github.com/maheswaradevo/wwx-backend/internal/authentications/delivery/http"
	authRepository "github.com/maheswaradevo/wwx-backend/internal/authentications/repository"
	authService "github.com/maheswaradevo/wwx-backend/internal/authentications/service"

	projectHTTPDelivery "github.com/maheswaradevo/wwx-backend/internal/projects/delivery/http"
	projectRepository "github.com/maheswaradevo/wwx-backend/internal/projects/repository"
	projectService "github.com/maheswaradevo/wwx-backend/internal/projects/service"

	"go.uber.org/zap"
)

func Init(router *echo.Echo, db *sql.DB, logger *zap.Logger, middleware echo.MiddlewareFunc) {
	appRestricted := router.Group("api/v1")
	appRestricted.Use(middleware)
	{

		InitProjectModule(appRestricted, db, logger)
	}
	app := router.Group("api/v1")
	{
		InitAuthModule(app, db, logger)
	}
}

func InitAuthModule(routerGroup *echo.Group, db *sql.DB, logger *zap.Logger) *echo.Group {
	authRepo := authRepository.NewAuthRepository(db, logger)
	authService := authService.NewAuthService(authRepo, logger)
	return authHTTPDelivery.AuthenticationNewDelivery(authService, routerGroup, logger)
}

func InitProjectModule(routerGroup *echo.Group, db *sql.DB, logger *zap.Logger) *echo.Group {
	projectRepo := projectRepository.NewProjectRepository(db, logger)
	projectService := projectService.NewProjetService(projectRepo, logger)
	return projectHTTPDelivery.ProjectNewDelivery(projectService, routerGroup, logger)
}
