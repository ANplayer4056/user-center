package handler

import (
	"fmt"
	"log"
	"net/http"
	"usercenter/app/internals/database"

	"github.com/gin-gonic/gin"
)

// DeleteUser ===>  刪除使用者的 api
func DeleteUser(c *gin.Context) {

	//  定義 api 接收的參數 struct
	type ReqUser struct {
		ID int `json:"id" binding:"required"`
	}

	//  取得 JSON data 參數
	req := ReqUser{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//  跟 DB 取得連線
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("DB connect failed ===> ", err)
	}

	//  處理 刪除 User
	if err = db.Delete(&req.ID).Error; err != nil {
		log.Printf("Error Message is %v ", err.Error())
		c.JSON(200, gin.H{
			"statusCode": 1004,
			"message":    "delete failed",
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
