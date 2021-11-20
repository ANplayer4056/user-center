package internal

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "root:example@tcp(127.0.0.1:3306)/backend_user?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//  gorm 連線池設定
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)                  // 最大的閒置連線數量
	sqlDB.SetMaxOpenConns(100)                 // 最大可存活連線數量，超過的必須等待
	sqlDB.SetConnMaxLifetime(time.Second * 10) // 閒置時間設定， 使用計算的方式 ex:time.Second * 10

	tx := db.Debug()

	return tx, err
}