package handler

import (
	"fmt"
	"golang_practice/app/internal"
	"log"
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

//  UpdateUser ===>  更新單一使用者的 api
func UpdateUser(c *gin.Context) {

	//  定義 api 接收的參數 struct
	type ReqUser struct {
		Username string `form:"Username" json:"Username" binding:"required"`
		Password string `form:"Password" json:"Password"`
	}

	//  取得 JSON data 參數
	reqParmams := ReqUser{}
	if err := c.ShouldBindJSON(&reqParmams); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	//  跟 DB 取得連線
	db, err := internal.ConnectDB()
	if err != nil {
		fmt.Println("DB connect failed ===> ", err)
	}

	//  處理 刪除 User
	//  structs.Map >> 当通过 struct 更新时，GORM 只会更新非零字段。 如果您想确保指定字段被更新，你应该使用 Select 更新选定字段，或使用 map 来完成更新操作
	mapDB := structs.Map(reqParmams.Password)
	if err = db.Model(&reqParmams.Password).Where("Username = ?", reqParmams.Username).Updates(&mapDB).Error; err != nil {
		log.Printf("Error Message is %v ", err.Error())

		c.JSON(200, gin.H{
			"statusCode": 1001,
			"message":    "update Error",
		})
		return
	}

	c.JSON(200, gin.H{
		"statusCode": 200,
		"message":    "update success",
	})
}
