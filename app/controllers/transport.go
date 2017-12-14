package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/zanjs/y-mugg-v3/app/models"
	"github.com/zanjs/y-mugg-v3/app/services"
)

// TransportController is 销量记录
type TransportController struct {
	Controller
}

// GetAll is get all Transports
func (ctl TransportController) GetAll(c echo.Context) error {

	var (
		datas       []models.Transport
		queryparams models.QueryParams
		page        models.PageModel
		err         error
	)
	queryparams = ctl.GetQueryParams(c)

	datas, page, err = services.TransportServices{}.GetAll(queryparams)
	if err != nil {
		return ctl.ResponseError(c, http.StatusForbidden, err.Error())
	}

	dataAll := echo.Map{
		"data": datas,
		"page": page,
	}

	return ctl.ResponseSuccess(c, dataAll)
}

// GetAllWhereTime is get all Transports
func (ctl TransportController) GetAllWhereTime(c echo.Context) error {

	var (
		sales       []models.Transport
		queryparams models.QueryParams
		err         error
	)
	queryparams = ctl.GetQueryParams(c)

	sales, err = services.TransportServices{}.WhereTime(queryparams)
	if err != nil {
		return ctl.ResponseError(c, http.StatusForbidden, err.Error())
	}

	return ctl.ResponseSuccess(c, sales)
}

// Update is update sales
func (ctl TransportController) Update(c echo.Context) error {
	// Parse the content
	data := new(models.Transport)

	quantity, _ := strconv.Atoi(c.FormValue("quantity"))

	data.Quantity = quantity

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetTransportByID(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// update record data
	err = m.UpdateTransport(data)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, m)
}
