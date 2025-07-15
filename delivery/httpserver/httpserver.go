package httpserver

import (
	"GameApp/config"
	"GameApp/services/userservice"
	// "github.com/labstack/echo/v4/middleware"



)

type Server struct {
	config config.Config
	userserv userservice.Service
}

func New(config config.Config,userserv userservice.Service) Server{
	return Server{
		config: config,
		userserv: userserv,
	}
}

// func (s Server) Serv(){
// 	e :=echo.


// }