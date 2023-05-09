package mysql

import (
	"fmt"
	"log"
	"os"
	"time"
	"u-future-api/util/exception"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbGlobal *gorm.DB

func InitDatabase() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,         // Don't include params in the SQL log
			Colorful:                  false,        // Disable color
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
	)
	if dbGlobal != nil {
		return dbGlobal
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
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
