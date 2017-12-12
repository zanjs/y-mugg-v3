package models

import (
	"time"

	"github.com/zanjs/y-mugg-v3/db"
)

type (
	// Inventory is 库存
	Inventory struct {
		BaseModel
		WareroomID int      `json:"wareroom_id"`
		ProductID  int      `json:"product_id"`
		Quantity   int      `json:"quantity"`
		Product    Product  `json:"product"`
		Wareroom   Wareroom `json:"wareroom"`
	}
)

// GetInventoryByID is
func GetInventoryByID(id int) (Inventory, error) {
	var (
		data Inventory
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

// UpdateInventory is
func (m *Inventory) UpdateInventory(data *Inventory) error {
	var err error

	m.UpdatedAt = time.Now()
	m.Quantity = data.Quantity
	m.ID = data.ID
	tx := gorm.MysqlConn().Begin()
	if err = tx.Save(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
