package model

import (
	"time"

	"gorm.io/gorm"
)

type UserList struct {
	ID        int    `gorm:"priamrykey"`
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
