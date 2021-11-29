package handler

import (
	"fmt"
	"log"
	"net/http"
	"usercenter/app/internals/database"
	"usercenter/app/model"

	"github.com/gin-gonic/gin"
)

//  Login ===> 使用者登入API
func Login(c *gin.Context) {

	//  定義 api 接收的參數 struct
	type UserLists struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	//  取得 JSON data 參數
	var json UserLists
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//  跟 DB 取得連線
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("DB connect failed ===> ", err)
	}

	//  處理 DB 查詢 (找到的資料) ===> 只有一筆
	dbaccept := model.UserList{} //   dbaccept

	//  找第一筆吻合資料
	result := db.Where("Username = ? AND Password  = ?", json.Account, json.Password).First(&dbaccept)

	// fmt.Printf("====>%+v \n", dbaccept)

	if result.Error != nil {
		log.Printf("Error Message is %v ", result.Error)
		c.JSON(200, gin.H{
			"statusCode": 1001,
			"message":    "login failed",
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

	obj := UserInfo{}
	obj.ID = dbaccept.ID
	obj.Username = dbaccept.Username
	obj.Status = dbaccept.Status
	obj.Depart = dbaccept.Depart
	obj.Level = dbaccept.Level

	c.JSON(200, gin.H{
		"statusCode": 0,
		"message":    "",
		"data":       obj,
	})
}
