package userhandler

import (
	"GameApp/param"
	"GameApp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) userLogin(c echo.Context) error {

	var req param.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)

	}

	if fildeerror ,err:=h.userValidator.ValidateLoginRequest(req);err!=nil{
		msg,code:=httpmsg.Error(err)
		return c.JSON(code,echo.Map{
			"message":msg,
			"errors":fildeerror,
		})
	}


	resp, err := h.userserv.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, resp)
}
