package handler

import (
	"fmt"
	"log"
	"net/http"
	"usercenter/app/internals/database"
	"usercenter/app/model"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

//  UpdatePassword ===>  更新使用者自己的密碼
func UpdatePassword(c *gin.Context) {

	//  定義 api 接收的參數 struct
	type UserInfo struct {
		Username      string `json:"username" binding:"required"`
		OldPassword   string `json:"oldpassword" binding:"required"`
		NewPassword   string `json:"newpassword" binding:"required"`
		CheckPassword string `json:"checkpassword" binding:"required"`
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

	//  查詢 user password
	dbaccept := model.UserList{} //   dbaccept
	result := db.Where("Username = ?", reqParmams.Username).First(&dbaccept)

	if result.Error != nil {
		log.Printf("Error Message is %v ", result.Error)
		c.JSON(200, gin.H{
			"statusCode": 1001,
			"message":    "用戶資料錯誤",
		})
		return
	}

	//  定義 更新DB struct
	type UpdateInfo struct {
		Username string
		Password string
	}

	//  判斷密碼正確?
	if dbaccept.Password != reqParmams.OldPassword {

		fmt.Println("Error ===>  密碼不符")
		c.JSON(200, gin.H{
			"statusCode": 1006,
			"message":    "密碼不符",
		})
		return
	}

	//  處理  更新 User Password
	mapDB := structs.Map(&UpdateInfo{})
	if err = db.Model(&UserInfo{}).Where("Username = ? ", reqParmams.Username).Updates(&mapDB).Error; err != nil {
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
