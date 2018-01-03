package models

import (
	"time"

	"github.com/zanjs/y-mugg-v3/db"
)

type (
	// Sale is 销售
	Sale struct {
		BaseModel
		WareroomID int      `json:"wareroom_id"`
		ProductID  int      `json:"product_id"`
		Quantity   int      `json:"quantity"`
		Product    Product  `json:"product,omitempty"`
		Wareroom   Wareroom `json:"wareroom,omitempty"`
	}

	// SaleOnly is
	SaleOnly struct {
		BaseModel
		WareroomID int `json:"wareroom_id"`
		ProductID  int `json:"product_id"`
		Quantity   int `json:"quantity"`
	}
)

// GetSaleByID is
func GetSaleByID(id uint64) (Sale, error) {
	var (
		data Sale
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Last(&data, id).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()

	return data, err
}

// UpdateSale is
func (m *Sale) UpdateSale(data *Sale) error {
	var err error

	m.UpdatedAt = time.Now()
	m.Quantity = data.Quantity

	tx := gorm.MysqlConn().Begin()
	if err = tx.Save(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
