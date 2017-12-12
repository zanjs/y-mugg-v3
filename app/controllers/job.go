package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/zanjs/y-mugg-v3/app/models"
	"github.com/zanjs/y-mugg-v3/app/queues"
	"github.com/zanjs/y-mugg-v3/app/utils"
)

// JobController is
type JobController struct {
	Controller
}

// SyncQnventory get all products and wareroom
func (ctl JobController) SyncQnventory(c echo.Context) error {
	var (
		data models.ProductWareroom
		err  error
	)
	data.Products, err = models.GetProducts()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	data.Warerooms, err = models.GetWarerooms()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	go func() {

		for _, v := range data.Warerooms {
			var wareroom models.Wareroom
			wareroom = v
			// numbering := wareroom.Numbering
			// fmt.Println(numbering)

			for _, v2 := range data.Products {

				var product models.Product
				product = v2
				// fmt.Println(product)

				var qmProduct models.QMProduct
				qmProduct.ItemCode = product.ExternalCode
				qmProduct.WarehouseCode = wareroom.Numbering
				qmProduct.OwnerCode = "bkyy"
				qmProduct.InventoryType = "ZP"

				var qmRequest models.QMRequest
				qmRequest = utils.Parameter("inventory.query", qmProduct)

				var inventory models.Inventory
				inventory.ProductID = product.ID
				inventory.WareroomID = wareroom.ID

				queues.QMHTTPPostV2(qmRequest, inventory)

			}

		}

	}()

	return c.JSON(http.StatusOK, data)
}

// SyncQnventoryV1 get all products and wareroom
func (ctl JobController) SyncQnventoryV1(c echo.Context) error {
	var (
		data models.ProductWareroom
		err  error
	)
	data.Products, err = models.GetProducts()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	data.Warerooms, err = models.GetWarerooms()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	go func() {

		for _, v := range data.Warerooms {
			var wareroom models.Wareroom
			wareroom = v
			numbering := wareroom.Numbering
			fmt.Println(numbering)

			for _, v2 := range data.Products {

				var product models.Product
				product = v2
				fmt.Println(product)

				var qmProduct models.QMProduct
				qmProduct.ItemCode = product.ExternalCode
				qmProduct.WarehouseCode = wareroom.Numbering
				qmProduct.OwnerCode = "bkyy"
				qmProduct.InventoryType = "ZP"

				var qmRequest models.QMRequest
				qmRequest = utils.Parameter("inventory.query", qmProduct)

				var record models.Record
				record.ProductID = product.ID
				record.WareroomID = wareroom.ID

				queues.QMHTTPPost(qmRequest, record)

			}

		}

	}()

	dataAll := echo.Map{
		"data":    data,
		"message": "同步数据中,稍等几分钟哦 ",
	}

	return c.JSON(http.StatusOK, dataAll)
}
