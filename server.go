package main

import (
	"fmt"
	"github.com/spf13/viper"
	. "github.com/trinhquocviet/template__go__api__echo/initialization"
)

func main() {
	err := LoadConfiguration()
	if err != nil {
		fmt.Println("An error occurred: ", err)
	} else {
		e := ConfigEchoFramework()
		InitRouter(e)

		// Server
		e.Logger.Fatal(e.Start(":" + viper.GetString("Port")))
	}
}
