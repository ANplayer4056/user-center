package handler

import (
	"fmt"
	"log"
	"net/http"
	"usercenter/app/internals/database"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

//  UpdatePassword ===>  更新使用者自己的密碼
func UpdatePassword(c *gin.Context) {

	//  定義 api 接收的參數 struct
	type UserInfo struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	//  取得 JSON data 參數
	reqParmams := UserInfo{}
	if err := c.ShouldBindJSON(&reqParmams); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	//  跟 DB 取得連線
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("DB connect failed ===> ", err)
	}

	//  處理  更新 User Password
	mapDB := structs.Map(&UserInfo{})
	if err = db.Model(&UserInfo{}).Where("Username = ?", reqParmams.Username).Updates(&mapDB).Error; err != nil {
		log.Printf("Error Message is %v ", err.Error())

		c.JSON(200, gin.H{
			"statusCode": 1005,
			"message":    "Update User Password Error",
			"data":       "",
		})
		return
	}

	c.JSON(200, gin.H{
		"statusCode": 0,
		"message":    "",
		"data":       "",
	})
}
