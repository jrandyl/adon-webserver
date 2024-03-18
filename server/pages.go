package server

import (
	"github.com/jrandyl/adon-webserver/client/src/pages/fragments"
	"github.com/jrandyl/adon-webserver/client/src/pages/templates"
	"github.com/labstack/echo/v4"
)

func (s *Server) index(c echo.Context) error {
	isRequestHX := c.Request().Header.Get("Hx-Request") != ""

	if isRequestHX {
		return s.Render(c, fragments.Index())
	}

	return s.Render(c, templates.Index())
}
