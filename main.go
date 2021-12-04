package main

import (
	"fmt"

	"usercenter/app/internals/database"
	"usercenter/router"

	"github.com/gin-gonic/gin"
)

//  call before main
func init() {
	if err := database.DBcheckTable(); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := database.DBcheckDepartTable(); err != nil {
		fmt.Println(err.Error())
		return
	}
}

/**
* gin :
* 	golang 所開發出來的 web framework
*	用來 建立 API
*
*
* gorm :
* 	是Golang語言中一款效能極好的ORM庫
*	用來與 db 溝通及處理資料表等使用
*
 */

// main ===> main.go 進入點
func main() {

	// defer 不管有無錯誤都會 run
	defer func() {
		if err := recover(); err != nil {
			// 補上將err傳至telegram
			fmt.Println("[❌ Fatal❌ ] HTTP:", err)
		}
	}()
	// 使用 gin 來製作 CRUD API
	r := gin.Default()
	router.BackendUser(r)
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
