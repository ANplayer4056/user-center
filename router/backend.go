package router

import (
	"usercenter/app/handler"

	"github.com/gin-gonic/gin"
)

func BackendUser(r *gin.Engine) {

	r.POST("/createUser", func(c *gin.Context) {
		handler.CreateUser(c)
	})
	r.DELETE("/deleteUser", handler.DeleteUser)
	r.PUT("/updateUser", handler.UpdateUser)
	r.POST("/queryUser", handler.QueryUser)
}
