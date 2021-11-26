package model

type DepartList struct {
	ID         int    `gorm:"priamrykey"`
	Departname string `gorm:"column:departname"`
	DepartCode string `gorm:"column:departcode"`
}
