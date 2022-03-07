package router

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/RaphaSalomao/gin-budget-control/controller"
	_ "github.com/RaphaSalomao/gin-budget-control/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	srvPort := fmt.Sprintf(":%s", os.Getenv("SRV_PORT"))
	// host := os.Getenv("SRV_HOST")

	router := gin.Default()

	unauthorized := router.Group("/budget-control/api/v1")
	{
		unauthorized.POST("/user", controller.CreateUser)
		unauthorized.POST("/user/authenticate", controller.Authenticate)
		unauthorized.GET("/health", controller.Health)
		unauthorized.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	receipt := router.Group("/budget-control/api/v1/receipt", authMiddleware())
	{
		receipt.POST("", controller.CreateReceipt)
		receipt.GET("", controller.FindAllReceipts)
		receipt.GET("/:id", controller.FindReceipt)
		receipt.PUT("/:id", controller.UpdateReceipt)
		receipt.DELETE("/:id", controller.DeleteReceipt)
		receipt.GET("/period/:year/:month", controller.ReceiptsByPeriod)
	}

	expense := router.Group("/budget-control/api/v1/expense", authMiddleware())
	{
		expense.POST("", controller.CreateExpense)
		expense.GET("", controller.FindAllExpenses)
		expense.GET("/:id", controller.FindExpense)
		expense.PUT("/:id", controller.UpdateExpense)
		expense.DELETE("/:id", controller.DeleteExpense)
		expense.GET("/period/:year/:month", controller.ExpensesByPeriod)
	}

	summary := router.Group("/budget-control/api/v1/summary", authMiddleware())
	{
		summary.GET("/:year/:month", controller.MonthBalanceSumary)
	}

	go router.Run(srvPort)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Server is running")
	<-quit
}
