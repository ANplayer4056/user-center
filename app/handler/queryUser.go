package handler

import (
	"fmt"
	"log"
	"net/http"
	"usercenter/app/internals/database"
	"usercenter/app/model"

	"github.com/gin-gonic/gin"
)

// QueryUser ===>  查詢指定用戶的 api
func QueryUser(c *gin.Context) {

	//  定義 api 接收的參數 struct
	type ReqUser struct {
		Username string `form:"Username" json:"Username" binding:"required"`
	}

	//  取得 JSON data 參數
	reqParmams := ReqUser{}
	if err := c.ShouldBindJSON(&reqParmams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//  跟 DB 取得連線
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("DB connect failed ===> ", err)
	}

	//  處理 DB 查詢
	dbaccept := []model.UserList{} //   dbaccept

	//  db.Find 可找到多個值
	result := db.Where("Username = ?", reqParmams.Username).Find(&dbaccept)

	if result.Error != nil {
		log.Printf("Error Message is %v ", result.Error)
		c.JSON(200, gin.H{
			"statusCode": 1001,
			"message":    "query failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"statusCode": 200,
		"message":    dbaccept,
	})
}
