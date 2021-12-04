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
	r.GET("/queryUser", handler.QueryUser)

	r.POST("/login", handler.Login)
	r.GET("/userlist", handler.UserList)
	r.PUT("/updatePassword", handler.UpdatePassword)
	r.PUT("/resetUserPassword", handler.ResetUserPassword)

	r.GET("/departlist", handler.DepartList)
	r.POST("/createDepart", handler.CreateDepart)
	r.DELETE("/deleteDepart", handler.DeleteDepart)
	r.PUT("/updateDepart", handler.UpdateDepart)

}
