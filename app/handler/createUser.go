package handler

import (
	"net/http"
	"usercenter/app/internals/database"
	"usercenter/app/model"

	"github.com/gin-gonic/gin"
)

//  CreateUser ===>  新增使用者的 api
//   定義 接收傳送過來的數據(params) / json:"API的欄位名稱" /  binding:"required" 必要欄位

func CreateUser(c *gin.Context) {

	//  定義 api 接收的參數 struct
	type UserInfo struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Status   bool   `json:"status" binding:"required"`
		Depart   string `json:"depart" binding:"required"`
		Level    int    `json:"level" binding:"required"`
	}

	//  取得 JSON data 參數
	var json UserInfo
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//  跟 DB 取得連線
	db, err := database.ConnectDB()
	if err != nil {

		c.JSON(200, gin.H{
			"statusCode": 500,
			"message":    "connect  failed " + err.Error(),
			"data":       "",
		})
		return
	}

	//  處理 新增 User
	if err = db.Model(&model.UserList{}).Create(map[string]interface{}{
		"username": json.Username, "password": json.Password, "status": json.Status, "depart": json.Depart, "level": json.Level,
	}).Error; err != nil {

		c.JSON(200, gin.H{
			"statusCode": 1002,
			"message":    "create failed " + err.Error(),
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
