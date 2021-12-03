package handler

import (
	"fmt"
	"net/http"
	"usercenter/app/internals/database"
	"usercenter/app/model"

	"github.com/gin-gonic/gin"
)

//  CreateDepart ===>  新增部門的 api
func CreateDepart(c *gin.Context) {

	//  定義 api 接收的參數 struct
	type DepartInfo struct {
		Departname string `json:"departName" binding:"required"`
		DepartCode string `json:"departCode" binding:"required"`
	}

	//  取得 JSON data 參數
	var json DepartInfo
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//  跟 DB 取得連線
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("DB connect failed ===> ", err)
	}

	//  處理 新增 User
	if err = db.Model(&model.DepartList{}).Create(map[string]interface{}{
		"departName": json.Departname, "departCode": json.DepartCode,
	}).Error; err != nil {

		c.JSON(200, gin.H{
			"statusCode": 1008,
			"message":    "create failed",
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
