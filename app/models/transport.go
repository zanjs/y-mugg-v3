package models

import (
	"time"

	"github.com/zanjs/y-mugg-v3/db"
)

type (
	// Transport is 托运记录
	Transport struct {
		BaseModel
		WareroomID int      `json:"wareroom_id"`
		ProductID  int      `json:"product_id"`
		Quantity   int      `json:"quantity"`
		Product    Product  `json:"product"`
		Wareroom   Wareroom `json:"wareroom"`
	}
)

// GetTransportByID is
func GetTransportByID(id uint64) (Transport, error) {
	var (
		data Transport
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

// UpdateTransport is
func (m *Transport) UpdateTransport(data *Transport) error {
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
