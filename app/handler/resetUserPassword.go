package handler

import (
	"fmt"
	"log"
	"net/http"
	"usercenter/app/internals/database"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

//  ResetUserPassword ===>  重設用戶密碼的 api
func ResetUserPassword(c *gin.Context) {

	//  定義 api 接收的參數 struct
	type ReqUser struct {
		ID       int    `json:"id" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	//  取得 JSON data 參數
	reqParmams := ReqUser{}
	if err := c.ShouldBindJSON(&reqParmams); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	//  跟 DB 取得連線
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("DB connect failed ===> ", err)
	}

	//  處理  更新 User
	//  structs.Map >> 当通过 struct 更新时，GORM 只会更新非零字段。 如果您想确保指定字段被更新，你应该使用 Select 更新选定字段，或使用 map 来完成更新操作
	mapDB := structs.Map(&ReqUser{})
	if err = db.Model(&ReqUser{}).Where("Username = ?", reqParmams.Username).Updates(&mapDB).Error; err != nil {
		log.Printf("Error Message is %v ", err.Error())

		c.JSON(200, gin.H{
			"statusCode": 1003,
			"message":    "update Error",
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
