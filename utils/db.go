package utils

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var DbErr error

func initDB() {
	dbConfig := GetConfig().Database
	dns := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	Db, DbErr = gorm.Open(mysql.Open(dns), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if DbErr != nil {
		Logger.Fatalln("数据库连接错误: ", DbErr)
	}

	sqlDb, sqlDbErr := Db.DB()
	if sqlDbErr != nil {
		Logger.Fatalln("数据库连接错误: ", sqlDbErr)
	}

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	// SetMaxOpenCons 设置数据库的最大连接数量。
	// SetConnMaxLifetime 设置连接的最大可复用时间。
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(10 * time.Second)
}
