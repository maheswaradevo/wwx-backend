package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/maheswaradevo/wwx-backend/pkg"
	"github.com/maheswaradevo/wwx-backend/pkg/config"
	"go.uber.org/zap"
)

func main() {
	config.Init()

	cfg := config.GetConfig()

	logger, _ := zap.NewProduction()

	db := config.GetDatabase(cfg.Database.Username, cfg.Database.Password, cfg.Database.Address, cfg.Database.Name)

	app := echo.New()

	pkg.Init(app, db, logger)
	app.Validator = nil

	address := fmt.Sprintf("%s:%s", "0.0.0.0", cfg.Port)

	app.Start(address)
}
