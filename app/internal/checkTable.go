package internal

import (
	"fmt"
	"golang_practice/app/model"
)

/*
DBcheckTable ===> deal with db AutoMigrate
                  處理 AutoMigrate
*/
func DBcheckTable() error {

	db, err := ConnectDB()
	if err != nil {
		fmt.Println("DB connect failed ===> ", err)
	}

	if err = db.AutoMigrate(&model.UserList{}); err != nil {
		fmt.Println("DB Migrate failed ===> ", err)
	}

	return err
}
