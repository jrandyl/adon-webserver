package server

import (
	"github.com/jrandyl/adon-webserver/client"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) SetupRouter() {

	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())
	s.router.Use(middleware.Secure())

	client.RegisterHandler(s.router)

	//pages
	s.router.GET("/", s.index)

	//components
	s.router.GET("/laptop", s.laptop)
	s.router.GET("/description", s.desciption)
	s.router.GET("/content", s.content)
	s.router.GET("/efficient", s.efficient)
	s.router.GET("/secure", s.secure)
	s.router.GET("/support", s.support)
	s.router.GET("/iphone", s.iphone)
}
