package queues

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zanjs/y-mugg-v3/app/models"
	"github.com/zanjs/y-mugg-v3/app/services"
	"github.com/zanjs/y-mugg-v3/app/utils"
)

// QMHTTPPostV2 is
func QMHTTPPostV2(qmRequest models.QMRequest, inventory models.Inventory) {

	// fmt.Println("qmRequest:", qmRequest)
	url := qmRequest.URL

	post := qmRequest.Body

	// fmt.Println("post:? ", post)

	// fmt.Println("URL:>", url)

	var xmlStr = []byte(post)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(xmlStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	// boydStr := string(body)

	qmResponse := models.QMResponse{}

	// body2 := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	// <response>
	// 	<flag>success</flag>
	// 	<code>SUCCESS</code>
	// 	<message>查询成功!</message>
	// 	<items>
	// 		<warehouseCode>B01</warehouseCode>
	// 		<itemCode>1371937585362246455</itemCode>
	// 		<itemId>1700046145</itemId>
	// 		<inventoryType>CC</inventoryType>
	// 		<quantity>3</quantity>
	// 		<lockQuantity>0</lockQuantity>
	// 	</items>
	// </response>`

	err = xml.Unmarshal([]byte(body), &qmResponse)
	if err != nil {
		fmt.Println("xml sp err :", err)
	}

	code := qmResponse.Code

	if code != "SUCCESS" {
		return
	}

	item := qmResponse.Items[0]

	// fmt.Println("xml Response Items :", item)
	fmt.Println("库存 Items :", item.Quantity)

	fmt.Println("记录商品信息 Items :", inventory)

	oldInventory, oErr := services.InventoryServices{}.GetByPId(inventory)
	// oldInventory, oErr := models.GetInventoryByID(inventory.ID)

	if oErr != nil {
		fmt.Println("查询旧数据：err:", oErr)
	}

	fmt.Println("查询旧数据：", oldInventory)

	var quantity = item.Quantity
	var oldQuantity = oldInventory.Quantity
	inventory.Quantity = quantity

	if oldInventory.ID == 0 {
		err := services.InventoryServices{}.Create(inventory)
		fmt.Println("旧库存 为空状态：", oldQuantity)
		fmt.Println("创建 newInventory 状态：", err)

		return
	}

	fmt.Println("旧库存ok吗：", oldQuantity)
	fmt.Println("新库存：", quantity)

	if oldQuantity == quantity {
		fmt.Println("不更新库存：", quantity)
		return
	}
	oldInventory.Quantity = quantity
	err = services.InventoryServices{}.Update(oldInventory)
	// data := new(models.Inventory)
	// data.Quantity = quantity
	// data.ID = oldInventory.ID
	// err = oldInventory.UpdateInventory(data)

	fmt.Println("更新库存 oldQuantity < quantity 状态：", err)
	if oldQuantity < quantity {
		fmt.Println("库存增加了：", err)
		return
	}

	fmt.Printf("进入销售减少状态")

	var salesCount = 0

	if oldQuantity > quantity {
		salesCount = oldQuantity - quantity
	}

	newSales := new(models.Sale)

	newSales.ProductID = inventory.ProductID
	newSales.WareroomID = inventory.WareroomID
	newSales.Quantity = salesCount

	services.SaleServices{}.Create(*newSales)
}

// QMHTTPPost is
func QMHTTPPost(qmRequest models.QMRequest, record models.Record) {

	fmt.Println("qmRequest:", qmRequest)
	url := qmRequest.URL

	post := qmRequest.Body

	fmt.Println("post:? ", post)

	fmt.Println("URL:>", url)

	var xmlStr = []byte(post)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(xmlStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	// boydStr := string(body)

	qmResponse := models.QMResponse{}

	// body2 := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	// <response>
	// 	<flag>success</flag>
	// 	<code>SUCCESS</code>
	// 	<message>查询成功!</message>
	// 	<items>
	// 		<warehouseCode>B01</warehouseCode>
	// 		<itemCode>1371937585362246455</itemCode>
	// 		<itemId>1700046145</itemId>
	// 		<inventoryType>CC</inventoryType>
	// 		<quantity>3</quantity>
	// 		<lockQuantity>0</lockQuantity>
	// 	</items>
	// </response>`

	err = xml.Unmarshal([]byte(body), &qmResponse)
	if err != nil {
		fmt.Println("xml sp err :", err)
	}

	code := qmResponse.Code

	if code != "SUCCESS" {
		return
	}

	item := qmResponse.Items[0]

	fmt.Println("xml Response Items :", item)
	fmt.Println("库存 Items :", item.Quantity)

	fmt.Println("记录商品信息 Items :", record)

	oldRecord, oErr := models.GetRecordLast(record.WareroomID, record.ProductID)

	if oErr != nil {
		fmt.Println("查询旧数据：err:", oErr)
	}

	fmt.Println("查询旧数据：", oldRecord)

	var initSales = 0
	var quantity = item.Quantity
	var oldQuantity = oldRecord.Quantity

	if utils.IsEmpty(oldQuantity) {
		fmt.Println("旧库存：", oldQuantity)
	}

	fmt.Println("旧库存：", oldQuantity)
	fmt.Println("新库存：", quantity)

	if oldQuantity == quantity {
		return
	}

	if oldQuantity > quantity {
		initSales = oldQuantity - quantity
	}

	newRecord := new(models.Record)

	newRecord.Quantity = item.Quantity
	newRecord.Sales = initSales
	newRecord.ProductID = record.ProductID
	newRecord.WareroomID = record.WareroomID

	models.CreateRecord(newRecord)
}
