package models

import (
	"github.com/zanjs/y-mugg-v3/db"
)

//CreateTable user
func CreateTable() error {
	gorm.MysqlConn().AutoMigrate(&User{}, &Article{}, &Product{}, &Wareroom{}, &Inventory{}, &Sale{})
	return nil
}
