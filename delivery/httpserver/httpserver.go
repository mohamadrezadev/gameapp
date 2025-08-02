package httpserver

import (
	"GameApp/config"
	"GameApp/services/authservice"
	"GameApp/services/userservice"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config config.Config
	authserv authservice.Service
	userserv userservice.Service

}

func New(config config.Config,userserv userservice.Service,authserv authservice.Service) Server{
	return Server{
		config: config,
		userserv: userserv,
		authserv: authserv,
	}
}

func (s Server) Serv(){
	e :=echo.New()

	// middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health-check",s.healthCheck)

	userGrop:=e.Group("/users")
	userGrop.GET("/profile",s.userProfile)
	userGrop.POST("/login",s.userLogin)
	userGrop.POST("/register",s.userRegister)


	//Start logger
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d",s.config.HTTPServer.Port)))


}