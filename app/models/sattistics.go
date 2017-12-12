package models

import "time"

type (
	// SattisticsProduct is
	SattisticsProduct struct {
		Product    Product     `json:"product"`
		Sattistics []Sattistic `json:"product_sattistics"`
	}
	// Sattistic is
	Sattistic struct {
		InventoryQuantity int       `json:"inventory_quantity"`
		SalesQuantity     int       `json:"sales_quantity"` // 数量
		Mean              int       `json:"mean"`
		CreatedAt         time.Time `json:"create_at"` // 平均值
	}
)
