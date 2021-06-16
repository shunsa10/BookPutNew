package main

import (
	"fmt"

	"github.com/shunsa10/BookPutNew/config"
)

func main()  {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)
	
}