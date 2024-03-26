package app

import (
	"fmt"
	"github.com/entl/test_task_middleware/internal/app/endpoint"
	"github.com/entl/test_task_middleware/internal/app/mw"
	"github.com/entl/test_task_middleware/internal/app/service"
	"github.com/labstack/echo/v4"
	"log"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()
	a.e = endpoint.New(a.s)

	a.echo = echo.New()
	a.echo.Use(mw.RoleCheck)

	a.echo.GET("/status", a.e.Status)
	a.echo.GET("/date", a.e.Date)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server is running...")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal("Failed to start server", err)
	}

	return nil
}
