package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type db struct {
	orm *gorm.DB
}

var DB *db

func New() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", "", "127.0.0.1", "3306", "cal_dev")

	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	DB = &db{orm: orm}
}
