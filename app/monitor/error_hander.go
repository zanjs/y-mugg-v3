package monitor

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/zanjs/y-mugg-v3/app/controllers"
)

// CustomHTTPErrorHandler is
func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	err = controllers.Controller{}.ResponseError(c, code, err.Error())
	c.Logger().Error(err)
}
