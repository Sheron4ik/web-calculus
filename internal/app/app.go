package app

import (
	"net/http"

	"github.com/Sheron4ik/web-calculus/internal/config"
	"github.com/Sheron4ik/web-calculus/internal/handlers"
)

type App struct {
	Config *config.Config
}

func New() *App {
	return &App{Config: config.New()}
}

func (a *App) Run() error {
	return http.ListenAndServe(":"+a.Config.Port, handlers.New())
}
