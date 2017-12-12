package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/zanjs/y-mugg-v3/app/models"
)

// WareroomController is
type WareroomController struct {
	Controller
}

// GetAll is get all warerooms
func (ctl WareroomController) GetAll(c echo.Context) error {
	var (
		warerooms []models.Wareroom
		err       error
	)
	warerooms, err = models.GetWarerooms()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, warerooms)
}

// Get is get one wareroom
func (ctl WareroomController) Get(c echo.Context) error {
	var (
		wareroom models.Wareroom
		err      error
	)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	wareroom, err = models.GetWareroomById(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, wareroom)
}

// Create is wareroom
func (ctl WareroomController) Create(c echo.Context) error {
	wareroom := new(models.Wareroom)
	wareroom.Title = c.FormValue("title")
	wareroom.Numbering = c.FormValue("numbering")

	err := models.CreateWareroom(wareroom)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusCreated, wareroom)
}

// Update is update wareroom
func (ctl WareroomController) Update(c echo.Context) error {
	// Parse the content
	wareroom := new(models.Wareroom)

	wareroom.Title = c.FormValue("title")
	wareroom.Numbering = c.FormValue("numbering")

	sortV := c.FormValue("sort")
	sort, _ := strconv.Atoi(sortV)

	wareroom.Sort = sort

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetWareroomById(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// update wareroom data
	err = m.UpdateWareroom(wareroom)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, m)
}

// Delete is wareroom
func (ctl WareroomController) Delete(c echo.Context) error {
	var err error

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetWareroomById(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	err = m.DeleteWareroom()
	return c.JSON(http.StatusNoContent, err)
}
