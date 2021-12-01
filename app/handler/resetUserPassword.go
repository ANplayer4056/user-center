package handler

import (
	"crypto/rand"
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

	//  定義 回傳 struct
	type ReSetPwd struct {
		ID       int
		Username string
		Password string
	}

	//  產生隨機密碼串
	n := 5 //  設定字母位數(A-Z)
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b) // 隨機密碼

	obj := ReSetPwd{}

	obj.ID = reqParmams.ID
	obj.Username = reqParmams.Username
	obj.Password = s

	//  處理  重設指定用戶密碼
	mapDB := structs.Map(&ReSetPwd{})
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
