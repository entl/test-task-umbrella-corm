package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, "Server healthy")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (e *Endpoint) Date(c echo.Context) error {
	d := e.s.DaysLeft()

	s := fmt.Sprintf("Days until 2025: %v", d)

	err := c.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}
