package testhelper

import "github.com/jinzhu/gorm"

func OpenDB() (*gorm.DB, error) {
	return gorm.Open("mysql", "notes:notes@(localhost)/notes?charset=utf8&parseTime=True&loc=Local")
}
