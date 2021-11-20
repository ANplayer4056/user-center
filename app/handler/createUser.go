package handler

import (
	"fmt"
	"golang_practice/app/internal"
	"golang_practice/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

//  CreateUser ===>  新增使用者的 api
//   定義 接收傳送過來的數據(params) / json:"API的欄位名稱" /  binding:"required" 必要欄位

func CreateUser(c *gin.Context) {

	//  定義 api 接收的參數 struct
	type UserLists struct {
		User     string `form:"user" json:"user" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	//  取得 JSON data 參數
	var json UserLists
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//  跟 DB 取得連線
	db, err := internal.ConnectDB()
	if err != nil {
		fmt.Println("DB connect failed ===> ", err)
	}

	//  處理 新增 User
	if err = db.Model(&model.UserList{}).Create(map[string]interface{}{
		"username": json.User, "password": json.Password,
	}).Error; err != nil {

		c.JSON(200, gin.H{
			"statusCode": 1001,
			"message":    "create failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"statusCode": 200,
		"userName":   json.User,
		"Password":   json.Password,
	})
}
