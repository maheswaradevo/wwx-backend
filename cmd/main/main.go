package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maheswaradevo/wwx-backend/pkg"
	"github.com/maheswaradevo/wwx-backend/pkg/config"
	"github.com/maheswaradevo/wwx-backend/pkg/middlewares"
	"go.uber.org/zap"
)

func main() {
	config.Init()

	cfg := config.GetConfig()

	logger, _ := zap.NewProduction()

	db := config.GetDatabase(cfg.Database.Username, cfg.Database.Password, cfg.Database.Address, cfg.Database.Port, cfg.Database.Name)

	app := echo.New()
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:5501", "https://main--resplendent-madeleine-f4afdb.netlify.app/"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	authMiddleware := middlewares.NewAuthMiddleware([]byte(cfg.ApiSecretKey))

	pkg.Init(app, db, logger, authMiddleware.AuthMiddleware())
	app.Validator = nil

	address := fmt.Sprintf("%s:%s", "0.0.0.0", cfg.Port)

	app.Start(address)
}
