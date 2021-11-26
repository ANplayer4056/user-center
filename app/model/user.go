package model

import (
	"time"

	"gorm.io/gorm"
)

type UserList struct {
	ID          int    `gorm:"priamrykey"`
	Username    string `gorm:"column:username"`
	Password    string `gorm:"column:password"`
	Status      bool   `gorm:"column:status"`
	DepartName  string `gorm:"column:departname"`
	AccountAuth int    `gorm:"column:auth"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
