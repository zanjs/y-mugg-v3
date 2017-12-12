package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/zanjs/y-mugg-v3/app/models"
	"github.com/zanjs/y-mugg-v3/app/services"
)

// SattisticsController is 统计
type SattisticsController struct {
	Controller
}

// WhereTime is time query
func (ctl SattisticsController) WhereTime(c echo.Context) error {
	var (
		products    []models.Product
		warerooms   []models.Wareroom
		queryparams models.QueryParams
		err         error
	)

	queryparams = ctl.GetQueryParams(c)

	products, err = models.GetProducts()
	if err != nil {
		return ctl.ResponseError(c, http.StatusForbidden, err.Error())
	}

	warerooms, err = models.GetWarerooms()
	if err != nil {
		return ctl.ResponseError(c, http.StatusForbidden, err.Error())
	}

	SattisticsProductsInt := []models.SattisticsProduct{}

	SattisticsProducts := SattisticsProductsInt[0:]

	for index := 0; index < len(products); index++ {
		var product models.Product
		product = products[index]

		pID := product.ID

		SattisticsInt := []models.Sattistic{}

		Sattistics := SattisticsInt[0:]

		for _, v2 := range warerooms {
			fmt.Println(v2)
			wID := v2.ID
			var sales []models.Sale
			var salesDay []models.Sale

			var qyps models.QueryParams

			qyps.StartTime = queryparams.StartTime
			qyps.EndTime = queryparams.EndTime
			qyps.ProductID = pID
			qyps.WareroomID = wID
			qyps.Day = queryparams.Day

			var inventory models.Inventory

			inventory.WareroomID = wID
			inventory.ProductID = pID

			inventory, err = services.InventoryServices{}.GetByPId(inventory)

			sales, err = services.SaleServices{}.WhereTime(qyps)
			salesDay, err = services.SaleServices{}.WhereDay(qyps)

			if err != nil {
				fmt.Println("查询 where sales time err: ", err, sales)
			}

			saleQuantity := 0
			for _, v3 := range sales {
				fmt.Println("v3: ", v3)
				saleQuantity = saleQuantity + v3.Quantity
			}

			saleMean := 0

			for _, v3 := range salesDay {
				fmt.Println("v3: ", v3)
				saleMean = saleMean + v3.Quantity
			}

			sattistic := models.Sattistic{}

			sattistic.SalesQuantity = saleQuantity
			sattistic.InventoryQuantity = inventory.Quantity
			sattistic.CreatedAt = inventory.CreatedAt
			sattistic.Mean = saleMean

			Sattistics = append(Sattistics, sattistic)
		}

		var sattisticsProduct models.SattisticsProduct
		sattisticsProduct.Product = product
		sattisticsProduct.Sattistics = Sattistics
		SattisticsProducts = append(SattisticsProducts, sattisticsProduct)
	}

	fmt.Println("SattisticsProducts:", SattisticsProducts)

	dataAll := echo.Map{
		"warerooms": warerooms,
		"products":  SattisticsProducts,
	}
	return ctl.ResponseSuccess(c, dataAll)
}
