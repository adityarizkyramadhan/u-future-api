package mysql

import (
	"fmt"
	"log"
	"os"
	"u-future-api/util/exception"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbGlobal *gorm.DB

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
	)
	if dbGlobal != nil {
		return dbGlobal
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("init db failed, %s\n", err)
	}
	dbGlobal = db
	return dbGlobal
}

func Migrate(model ...interface{}) error {
	if dbGlobal != nil {
		return dbGlobal.AutoMigrate(model...)
	}
	return exception.ErrDatabaseNull
}
