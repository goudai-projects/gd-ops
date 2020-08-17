package db

import (
	"github.com/goudai-projects/gd-ops/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(dsn string) {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("connect to database fail")
	}
}
