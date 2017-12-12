package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/zanjs/y-mugg-v3/app/models"
	"github.com/zanjs/y-mugg-v3/app/services"
)

// InventoryController is
type InventoryController struct {
	Controller
}

// GetAll is get all Sales
func (ctl InventoryController) GetAll(c echo.Context) error {

	var (
		datas       []models.Inventory
		page        models.PageModel
		queryparams models.QueryParams
		err         error
	)
	queryparams = ctl.GetQueryParams(c)

	datas, page, err = services.InventoryServices{}.GetAll(queryparams)
	if err != nil {
		return ctl.ResponseError(c, http.StatusForbidden, err.Error())
	}

	dataAll := echo.Map{
		"data": datas,
		"page": page,
	}

	return ctl.ResponseSuccess(c, dataAll)
}
