package context

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DataRes struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorRes struct {
	Status string      `json:"status"`
	ErrMsg interface{} `json:"data"`
}

func RetData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, DataRes{
		Status: "200",
		Data:   data,
	})
}

func RetError(c echo.Context, code int, status, errMsg string) error {
	return c.JSON(code, ErrorRes{
		Status: status,
		ErrMsg: errMsg,
	})
}
