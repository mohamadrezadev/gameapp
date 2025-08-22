package httpserver

import (
	"GameApp/config"
	"GameApp/services/authservice"
	"GameApp/services/userservice"
	"GameApp/validator/uservalidator"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config config.Config
	authserv authservice.Service
	userserv userservice.Service
	userValidator uservalidator.Validator

}

func New(config config.Config,userserv userservice.Service,authserv authservice.Service,uservalidator uservalidator.Validator) Server{
	return Server{
		config: config,
		userserv: userserv,
		authserv: authserv,
		userValidator: uservalidator,
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