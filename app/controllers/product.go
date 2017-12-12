package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/zanjs/y-mugg-v3/app/models"
)

// ProductController is
type ProductController struct {
	Controller
}

// GetAll is all products
func (ctl ProductController) GetAll(c echo.Context) error {
	var (
		products []models.Product
		err      error
	)
	products, err = models.GetProducts()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, products)
}

// Get is one product
func (ctl ProductController) Get(c echo.Context) error {
	var (
		product models.Product
		err     error
	)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	product, err = models.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, product)
}

// Create is create product
func (ctl ProductController) Create(c echo.Context) error {

	product := new(models.Product)

	product.Title = c.FormValue("title")
	product.ExternalCode = c.FormValue("external_code")
	sortV := c.FormValue("sort")
	sort, _ := strconv.Atoi(sortV)
	fmt.Println(sort)
	box, _ := strconv.Atoi(c.FormValue("box"))
	exceed, _ := strconv.Atoi(c.FormValue("exceed"))
	product.Sort = sort
	product.Box = box
	product.Exceed = exceed

	err := models.CreateProduct(product)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusCreated, product)
}

// Update is update product
func (ctl ProductController) Update(c echo.Context) error {
	// Parse the content
	product := new(models.Product)

	product.Title = c.FormValue("title")
	product.ExternalCode = c.FormValue("external_code")

	sortV := c.FormValue("sort")
	sort, _ := strconv.Atoi(sortV)

	box, _ := strconv.Atoi(c.FormValue("box"))
	exceed, _ := strconv.Atoi(c.FormValue("exceed"))

	product.Sort = sort
	product.Box = box
	product.Exceed = exceed

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// update product data
	err = m.UpdateProduct(product)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, m)
}

//Delete is product
func (ctl ProductController) Delete(c echo.Context) error {
	var err error

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	err = m.DeleteProduct()
	return c.JSON(http.StatusNoContent, err)
}
