package main

import (
	"fmt"
	"strings"

	"github.com/RaphaSalomao/gin-budget-control/database"
	"github.com/RaphaSalomao/gin-budget-control/router"
	"github.com/joho/godotenv"
)

// @title     Alura Backend Challenge 2nd Edition API
// @version   1.3.1
// @host      localhost:5000
// @BasePath  /
func main() {
	err := godotenv.Load()
	if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		fmt.Println("Error loading .env file, using default environment variables")
	}
	err = database.Connect()
	if err != nil {
		panic(err)
	}
	router.HandleRequests()
}
