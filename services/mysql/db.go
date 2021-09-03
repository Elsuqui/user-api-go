package mysql

import (
	"UserRestApi/helpers"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connnect() (*gorm.DB, error) {
	dbUser := helpers.GetEnvParamDefault("DB_USER", "root")
	dbPass := helpers.GetEnvParam("DB_PASS")
	dbHost := helpers.GetEnvParamDefault("DB_HOST", "127.0.0.1")
	dbPort := helpers.GetEnvParamDefault("DB_PORT", "3306")
	dbDatabase := helpers.GetEnvParam("DB_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            false,
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	helpers.HandleError(err)
	DB = db
	return db, err
}

func Close(db *gorm.DB) {
	dbInstance, err := db.DB()
	helpers.HandleError(err)
	if err != nil {
		defer dbInstance.Close()
	}
}
