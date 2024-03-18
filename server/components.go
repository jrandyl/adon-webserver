package server

import (
	"github.com/jrandyl/adon-webserver/client/src/components"
	"github.com/labstack/echo/v4"
)

func (s *Server) laptop(c echo.Context) error {
	return s.Render(c, components.Laptop())
}

func (s *Server) desciption(c echo.Context) error {
	return s.Render(c, components.Description())
}

func (s *Server) content(c echo.Context) error {
	return s.Render(c, components.Content())
}

func (s *Server) efficient(c echo.Context) error {
	return s.Render(c, components.Efficient())
}

func (s *Server) secure(c echo.Context) error {
	return s.Render(c, components.Secure())
}

func (s *Server) support(c echo.Context) error {
	return s.Render(c, components.Support())
}

func (s *Server) iphone(c echo.Context) error {
	return s.Render(c, components.Iphone())
}
