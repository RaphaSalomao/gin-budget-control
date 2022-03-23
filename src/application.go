package main

import (
	"github.com/RaphaSalomao/gin-budget-control/config"
	"github.com/RaphaSalomao/gin-budget-control/database"
	"github.com/RaphaSalomao/gin-budget-control/router"
)

// @title     Alura Backend Challenge 2nd Edition API
// @version   1.3.1
// @host      localhost:5000
// @BasePath  /
func main() {
	config.LoadEnvironment()
	err := database.Connect()
	if err != nil {
		panic(err)
	}
	router.RunServer()
}
