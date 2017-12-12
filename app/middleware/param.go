package middleware

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo"
	"github.com/zanjs/y-mugg-v3/app/models"
)

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Bear/6.0")
		c.Response().Header().Set("copy", "zanjs")
		return next(c)
	}
}

// QueryParam middleware adds a `queryparams` header to the request.
func QueryParam(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			queryparams models.QueryParams
		)

		qps := c.QueryParams()

		limitq := c.QueryParam("limit")
		offsetq := c.QueryParam("offset")
		startTimeq := c.QueryParam("start_time")
		endTime := c.QueryParam("end_time")

		limit, _ := strconv.Atoi(limitq)
		offset, _ := strconv.Atoi(offsetq)
		fmt.Println(qps)
		fmt.Println(limit)
		fmt.Println(offset)

		if limit == 0 {
			limit = 10
		}

		queryparams.Limit = limit
		queryparams.Offset = offset
		queryparams.StartTime = startTimeq
		queryparams.EndTime = endTime

		c.Set("queryparams", queryparams)

		return next(c)
	}
}
