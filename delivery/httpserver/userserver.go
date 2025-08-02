package httpserver

import (
	"GameApp/services/userservice"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s Server) userRegister(c echo.Context) error {
	var uReq userservice.RegisterRequest

	if err := c.Bind(&uReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	resp, err := s.userserv.Register(uReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, resp)
}

func (s Server) userLogin(c echo.Context) error {

	var req userservice.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)

	}
	resp, err := s.userserv.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, resp)
}

func (s Server) userProfile(c echo.Context) error {
	fmt.Println("c.GetAuthorization", c.Get("Authorization"))
	authtoken := c.Request().Header.Get("Authorization")
	claims, err := s.authserv.ParseToken(authtoken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())

	}
	resp, err := s.userserv.Profile(userservice.ProfileRequest{UserId: claims.UserID})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, resp)

}
