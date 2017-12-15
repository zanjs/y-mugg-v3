package queues

import (
	"fmt"

	"github.com/zanjs/y-mugg-v3/app/models"
	"github.com/zanjs/y-mugg-v3/app/services"
)

// ISSale is
func ISSale(product models.Product, inventory models.Inventory, salesCount int) {

	newSales := new(models.Sale)

	newSales.ProductID = inventory.ProductID
	newSales.WareroomID = inventory.WareroomID
	newSales.Quantity = salesCount

	box := product.Box
	exceed := product.Exceed

	remainder := 0

	isTtansport := false

	fmt.Println("product")
	fmt.Println(product)

	if box != 0 && exceed != 0 {
		if salesCount >= exceed {
			isTtansport = true
			remainder = salesCount % box
			newSales.Quantity = remainder
		}
	}

	transport := new(models.Transport)
	transport.ProductID = inventory.ProductID
	transport.WareroomID = inventory.WareroomID

	if isTtansport {
		fmt.Println("有托运")
		transport.Quantity = salesCount - remainder
	}

	if transport.Quantity != 0 {
		services.TransportServices{}.Create(*transport)
	}

	if newSales.Quantity == 0 {
		return
	}

	services.SaleServices{}.Create(*newSales)
}
