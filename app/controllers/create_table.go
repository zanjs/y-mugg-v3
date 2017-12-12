package controllers

import (
	"github.com/labstack/echo"
	"github.com/zanjs/y-mugg-v3/app/models"
)

//CreateTable is 同步表结构
func CreateTable(c echo.Context) error {
	err := models.CreateTable()
	if err != nil {
		panic(err)
	}

	return c.String(200, "createTable ok")
}
