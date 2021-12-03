package handler

import (
	"fmt"
	"log"
	"usercenter/app/internals/database"
	"usercenter/app/model"

	"github.com/gin-gonic/gin"
)

//  DepartList ===> 部門列表 API
func DepartList(c *gin.Context) {

	//  跟 DB 取得連線
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("DB connect failed ===> ", err)
	}

	//  處理 DB 查詢 (找到的資料) ===> 不只一筆
	dbaccept := []model.DepartList{} //   dbaccept

	//  sql 處理
	result := db.Find(&dbaccept)

	if result.Error != nil {
		log.Printf("Error Message is %v ", result.Error)
		c.JSON(200, gin.H{
			"statusCode": 1007,
			"message":    "find depart error",
			"data":       "",
		})
		return
	}

	//  回傳的資料集
	dbArr := []model.DepartList{}

	//  每一筆資料架構
	obj := model.DepartList{}

	for _, item := range dbaccept {
		obj.ID = item.ID
		obj.DepartCode = item.DepartCode
		obj.Departname = item.Departname
	}

	//  回傳組合好的資料
	c.JSON(200, gin.H{
		"statusCode": 0,
		"message":    "",
		"data":       dbArr,
	})
}
