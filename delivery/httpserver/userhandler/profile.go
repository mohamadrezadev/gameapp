package userhandler

import (
	"GameApp/param"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) userProfile(c echo.Context) error {
	fmt.Println("c.GetAuthorization", c.Get("Authorization"))
	authtoken := c.Request().Header.Get("Authorization")
	claims, err := h.authserv.ParseToken(authtoken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())

	}
	resp, err := h.userserv.Profile(param.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, resp)

}
