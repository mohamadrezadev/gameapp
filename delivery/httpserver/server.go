package httpserver

import (
	"GameApp/config"
	"GameApp/delivery/httpserver/userhandler"
	"GameApp/services/authservice"
	"GameApp/services/userservice"
	"GameApp/validator/uservalidator"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config      config.Config
	userhandler userhandler.Handler
}

func New(config config.Config, userserv userservice.Service, authserv authservice.Service, uservalidator uservalidator.Validator) Server {
	return Server{
		config:      config,
		userhandler: userhandler.New(authserv, userserv, uservalidator),
	}
}

func (s Server) Serv() {
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health-check", s.healthCheck)

	//Start logger
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))

}
