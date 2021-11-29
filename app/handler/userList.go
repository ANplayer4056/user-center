package handler

import (
	"fmt"
	"log"
	"usercenter/app/internals/database"
	"usercenter/app/model"

	"github.com/gin-gonic/gin"
)

//  UserList ===> 用戶列表 API
func UserList(c *gin.Context) {

	//  跟 DB 取得連線
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("DB connect failed ===> ", err)
	}

	//  處理 DB 查詢 (找到的資料) ===> 不只一筆
	dbaccept := []model.UserList{} //   dbaccept

	//  sql 處理
	result := db.Find(&dbaccept)

	if result.Error != nil {
		log.Printf("Error Message is %v ", result.Error)
		c.JSON(200, gin.H{
			"statusCode": 1002,
			"message":    "find users error",
			"data":       "",
		})
		return
	}

	//  定義回傳 struct
	type UserInfo struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Status   bool   `json:"status"`
		Depart   string `json:"depart"`
		Level    int    `json:"level"`
	}

	//  回傳的資料集
	dbArr := []UserInfo{}

	//  每一筆資料架構
	obj := UserInfo{}

	for _, item := range dbaccept {

		obj.ID = item.ID
		obj.Username = item.Username
		obj.Status = item.Status
		obj.Depart = item.Depart
		obj.Level = item.Level

		dbArr = append(dbArr, obj)
	}

	//  回傳組合好的資料
	c.JSON(200, gin.H{
		"statusCode": 0,
		"message":    "",
		"data":       dbArr,
	})
}
