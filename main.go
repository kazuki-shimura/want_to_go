package main

import (
	"fmt"
	"log"
	"want_to_go/config"
)

func main() {
	fmt.Println("Want to Go App Start")
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")
}
