package services

import (
	"fmt"
	"time"

	"github.com/zanjs/y-mugg-v3/app/middleware"
	"github.com/zanjs/y-mugg-v3/app/models"
	"github.com/zanjs/y-mugg-v3/db"
)

type (
	// TransportServices is
	TransportServices struct{}
)

// GetAll is
func (sev TransportServices) GetAll(q models.QueryParams) ([]models.Transport, models.PageModel, error) {
	var (
		datas []models.Transport
		page  models.PageModel
		// queryParams models.QueryParams
		err error
	)

	page.Limit = q.Limit
	page.Offset = q.Offset

	pID := q.ProductID
	wID := q.WareroomID

	tx := gorm.MysqlConn().Begin()

	if page.Offset == 0 {

		if pID == 0 && wID == 0 {
			err = tx.Preload("Wareroom").Preload("Product").Order("id desc").Limit(page.Limit).Find(&datas).Count(&page.Count).Error
		} else if pID != 0 && wID != 0 {
			err = tx.Where("product_id = ? AND wareroom_id = ?", pID, wID).Preload("Wareroom").Preload("Product").Order("id desc").Limit(page.Limit).Find(&datas).Count(&page.Count).Error
		} else if pID != 0 {
			err = tx.Where("product_id = ?", pID).Preload("Wareroom").Preload("Product").Order("id desc").Limit(page.Limit).Find(&datas).Count(&page.Count).Error
		} else if wID != 0 {
			err = tx.Where("wareroom_id = ?", wID).Preload("Wareroom").Preload("Product").Order("id desc").Limit(page.Limit).Find(&datas).Count(&page.Count).Error
		}

	} else {

		if pID == 0 && wID == 0 {
			err = tx.Preload("Wareroom").Preload("Product").Order("id desc").Limit(page.Limit).Find(&datas).Error
		} else if pID != 0 && wID != 0 {
			err = tx.Where("product_id = ? AND wareroom_id = ?", pID, wID).Preload("Wareroom").Preload("Product").Order("id desc").Limit(page.Limit).Find(&datas).Error
		} else if pID != 0 {
			err = tx.Where("product_id = ?", pID).Preload("Wareroom").Preload("Product").Order("id desc").Limit(page.Limit).Find(&datas).Error
		} else if wID != 0 {
			err = tx.Where("wareroom_id = ?", wID).Preload("Wareroom").Preload("Product").Order("id desc").Limit(page.Limit).Find(&datas).Error
		}

	}

	if err != nil {
		tx.Rollback()
		return datas, page, err
	}

	tx.Commit()

	return datas, page, err
}

// Create is
func (sev TransportServices) Create(m models.Transport) error {
	var err error

	m.CreatedAt = time.Now()
	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// Delete is
func (sev TransportServices) Delete(m models.Transport) error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// WhereTime is
func (sev TransportServices) WhereTime(q models.QueryParams) ([]models.Transport, error) {
	var (
		datas []models.Transport
		err   error
	)

	queryTime := middleware.QueryStartEndTime(q)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Order("id desc").Where("created_at BETWEEN ? AND ?", queryTime.StartTime, queryTime.EndTime).Where("wareroom_id = ? AND product_id = ?", q.WareroomID, q.ProductID).Find(&datas).Error; err != nil {
		tx.Rollback()
		return datas, err
	}
	tx.Commit()

	fmt.Println("datas:", datas)

	return datas, err
}

// WhereDay is 周期平均值
func (sev TransportServices) WhereDay(q models.QueryParams) ([]models.Transport, error) {
	var (
		datas []models.Transport
		err   error
	)

	queryTime := middleware.QueryStartDay(q)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Order("id desc").Where("created_at BETWEEN ? AND ?", queryTime.StartTime, queryTime.EndTime).Where("wareroom_id = ? AND product_id = ?", q.WareroomID, q.ProductID).Find(&datas).Error; err != nil {
		tx.Rollback()
		return datas, err
	}
	tx.Commit()

	fmt.Println("datas:", datas)

	return datas, err
}
