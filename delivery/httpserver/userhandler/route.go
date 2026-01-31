package userhandler

import (
	"GameApp/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	userGrop:=e.Group("/users")
	userGrop.GET("/profile",h.userProfile,middleware.Auth(h.authserv,h.authconfig))
	userGrop.POST("/login",h.userLogin)
	userGrop.POST("/register",h.userRegister)

}