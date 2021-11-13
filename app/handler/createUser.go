package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//  CreateUser ===>  新增使用者的 api
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

}
