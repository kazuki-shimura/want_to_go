package main

import (
	"fmt"
	"want_to_go/app/models"
)

func main() {
	// fmt.Println("Want to Go App Start")
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	/*Userの作成*/
	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@gmail.com"
	// u.Password = "testtest"
	// fmt.Println(u)
	// u.CreateUser()

	/*　IDを使用したUserの取得　*/
	u, _ := models.GetUser(1)
	fmt.Println(u)
}
