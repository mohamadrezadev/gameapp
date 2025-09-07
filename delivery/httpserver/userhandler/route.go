package userhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(e *echo.Echo) {
	userGrop:=e.Group("/users")
	userGrop.GET("/profile",h.userProfile)
	userGrop.POST("/login",h.userLogin)
	userGrop.POST("/register",h.userRegister)

}