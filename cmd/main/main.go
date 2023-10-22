package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/maheswaradevo/wwx-backend/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maheswaradevo/wwx-backend/pkg"
	"github.com/maheswaradevo/wwx-backend/pkg/config"
	"github.com/maheswaradevo/wwx-backend/pkg/middlewares"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
)

// @title Wonderworxs Dashboard API
// @version 1.0
// @description API used by Wonderworxs to manage projects from client
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email pundadevo21@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/license/LICENSE-2.0.html
// @host xxx-backend.deveureka.com
// @BasePath /

func main() {
	config.Init()

	cfg := config.GetConfig()

	logger, _ := zap.NewProduction()

	db := config.GetDatabase(cfg.Database.Username, cfg.Database.Password, cfg.Database.Address, cfg.Database.Port, cfg.Database.Name)

	app := echo.New()
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:5501", "https://wwx-dashboard.netlify.app", "https://wwx-dashboard.vercel.app", "https://xxx-backend.deveureka.com"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	authMiddleware := middlewares.NewAuthMiddleware([]byte(cfg.ApiSecretKey))

	pkg.Init(app, db, logger, authMiddleware.AuthMiddleware())
	app.Validator = nil
	app.GET("/swagger/*", echoSwagger.WrapHandler)
	address := fmt.Sprintf("%s:%s", "0.0.0.0", cfg.Port)

	app.Start(address)
}
