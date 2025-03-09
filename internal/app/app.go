package app

import (
	"fmt"

	"github.com/Sheron4ik/web-calculus/internal/config"
	"github.com/Sheron4ik/web-calculus/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	Config *config.Config
}

func New() *App {
	return &App{Config: config.New()}
}

func (a *App) Run() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.POST("/api/v1/calculate", handlers.HandleCalculate)
	e.GET("/api/v1/expressions", handlers.HandleListExpressions)
	e.GET("/api/v1/expressions/:id", handlers.HandleGetExpression)
	e.GET("/internal/task", handlers.HandleGetTask)
	e.POST("/internal/task", handlers.HandleUpdateTask)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", a.Config.Port)))
}
