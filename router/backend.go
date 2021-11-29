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

	r.POST("/login", handler.Login)
	r.GET("/userlist", handler.UserList)
	r.PUT("/updatePassword", handler.UpdatePassword)
	r.PUT("/resetUserPassword", handler.ResetUserPassword)
}
