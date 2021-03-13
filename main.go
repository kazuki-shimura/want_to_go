package main

import (
	"fmt"
	"want_to_go/config"
)

func main() {
	fmt.Println("Want to Go App Start")
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)
}
