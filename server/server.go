package server

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Server struct {
	router *echo.Echo
}

func New() (*Server, error) {

	server := &Server{
		router: echo.New(),
	}

	server.SetupRouter()

	return server, nil
}

func (s *Server) Start(port string) error {
	return s.router.Start(port)
}

func (s *Server) Render(c echo.Context, templ templ.Component) error {
	return templ.Render(c.Request().Context(), c.Response())
}
