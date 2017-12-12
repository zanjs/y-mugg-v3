package controllers

import (
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/zanjs/y-mugg-v3/app/models"
)

// ResourceController is
type ResourceController interface {
	GetAll(c echo.Context)
	Get(c echo.Context)
	Create(c echo.Context)
	Update(c echo.Context)
	Delete(c echo.Context)
}

type httpError struct {
	code    int
	Key     string `json:"error"`
	Message string `json:"message"`
}

// UnityJSON is Return to unity JSON Data
type UnityJSON struct {
	Status  string      `json:"status"`
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    string      `json:"code"`
}

func newHTTPError(code int, key string, msg string) *httpError {
	return &httpError{
		code:    code,
		Key:     key,
		Message: msg,
	}
}

// Error makes it compatible with `error` interface.
func (e *httpError) Error() string {
	return e.Key + ": " + e.Message
}

// httpErrorHandler customize echo's HTTP error handler.
func httpErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		key  = "ServerError"
		msg  string
	)

	if he, ok := err.(*httpError); ok {
		code = he.code
		key = he.Key
		msg = he.Message
	} else if ee, ok := err.(*echo.HTTPError); ok {
		code = ee.Code
		key = http.StatusText(code)
		msg = key
		// } else if config.Debug {
		// 	msg = err.Error()
	} else {
		msg = http.StatusText(code)
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			err := c.JSON(code, newHTTPError(code, key, msg))
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}

// Controller is base controller
type Controller struct{}

func (ctl Controller) buildSuccessData(data interface{}) map[string]interface{} {
	return echo.Map{
		"status": "success",
		"error":  false,
		"data":   data,
	}
}

func (ctl Controller) buildErrorData(message string) map[string]interface{} {
	return echo.Map{
		"status":  "error",
		"error":   true,
		"message": message,
	}
}

// ResponseSuccess is
func (ctl Controller) ResponseSuccess(c echo.Context, data interface{}) error {
	return c.JSON(200, ctl.buildSuccessData(data))
}

// ResponseError is
func (ctl Controller) ResponseError(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, ctl.buildErrorData(message))
}

// GetQueryParams is get queryParams
func (ctl Controller) GetQueryParams(c echo.Context) models.QueryParams {
	var (
		queryparams models.QueryParams
	)

	startTimeq := c.QueryParam("start_time")
	endTime := c.QueryParam("end_time")

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	wareroomID, _ := strconv.Atoi(c.QueryParam("wareroom_id"))
	productID, _ := strconv.Atoi(c.QueryParam("product_id"))
	day, _ := strconv.Atoi(c.QueryParam("day"))

	if limit == 0 {
		limit = 10
	}

	queryparams.Limit = limit
	queryparams.Offset = offset
	queryparams.StartTime = startTimeq
	queryparams.EndTime = endTime
	queryparams.ProductID = productID
	queryparams.WareroomID = wareroomID
	queryparams.Day = day
	return queryparams
}

// GetPathParam is
func (ctl Controller) GetPathParam(c echo.Context) models.PathParams {

	var (
		pathparams models.PathParams
	)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	pathparams.ID = id
	return pathparams
}

// GetUser is get user 获取用户
func (ctl Controller) GetUser(c echo.Context) models.User {
	var (
		user models.User
	)
	userJwt := c.Get("user").(*jwt.Token)
	claims := userJwt.Claims.(jwt.MapClaims)
	userID := int(claims["id"].(float64))
	user.ID = userID
	return user
}
