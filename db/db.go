package db

import "github.com/fitrah-firdaus/notes/config"
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Init() {
	conf := config.GetConfig()
	username := conf.GetString("db.mysql.username")
	password := conf.GetString("db.mysql.password")
	hostname := conf.GetString("db.mysql.hostname")
	dbname := conf.GetString("db.mysql.dbname")

	db, err := gorm.Open("mysql", username+":"+password+"@("+hostname+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
}

func GetDB() *gorm.DB {
	return db
}
