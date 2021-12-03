package handler

import (
	"fmt"
	"log"
	"net/http"
	"usercenter/app/internals/database"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

//  UpdateDepart ===>  更新單一部門的 api
func UpdateDepart(c *gin.Context) {

	//  定義 api 接收的參數 struct
	type ReqDepart struct {
		ID         int    `json:"id" binding:"required"`
		Departname string `json:"departname" binding:"required"`
		DepartCode string `json:"departCode" binding:"required"`
	}

	//  取得 JSON data 參數
	reqParmams := ReqDepart{}
	if err := c.ShouldBindJSON(&reqParmams); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	//  跟 DB 取得連線
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("DB connect failed ===> ", err)
	}

	//  處理  更新 部門
	mapDB := structs.Map(&ReqDepart{})
	if err = db.Model(&ReqDepart{}).Where("ID = ?", reqParmams.ID).Updates(&mapDB).Error; err != nil {
		log.Printf("Error Message is %v ", err.Error())

		c.JSON(200, gin.H{
			"statusCode": 1010,
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
