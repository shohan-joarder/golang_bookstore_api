package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect()  {
	dbUser := "root"
	dbPassword := ""
	dbDatabase := "book_store"
	d, err:= gorm.Open("mysql", dbUser+":"+dbPassword+"@/"+dbDatabase)

	if err != nil{
		panic(err.Error())
	}
	db = d
}

func GetDb() *gorm.DB{
	return db;
}