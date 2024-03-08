package bootstrap

import (
	"fmt"
	"go-api-template/logs"

	"gorm.io/driver/mysql"
	// "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Conn *gorm.DB

// func NewDatabaseConnection(env *Env) *gorm.DB {
// 	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable",
// 		env.Database.Username,
// 		url.QueryEscape(env.Database.Password),
// 		env.Database.DBHost,
// 		env.Database.DBPort,
// 		env.Database.DBName,
// 	)
// 	// Set up a custom pool configuration
// 	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("can not connect to database")
// 	}
// 	logs.Info("database connection success")
// 	return db
// }

func NewDatabaseDpiConnection(env *Env) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DatabaseDpi.Username,
		env.DatabaseDpi.Password,
		env.DatabaseDpi.DBHost,
		env.DatabaseDpi.DBPort,
		env.DatabaseDpi.DBName,
	)
	// fmt.Println(dsn)
	// Set up a custom pool configuration
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can not connect to database")
	}
	logs.Info("database connection success")

	Conn = db
	return db
}
