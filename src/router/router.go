package router

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/RaphaSalomao/gin-budget-control/controller"
	_ "github.com/RaphaSalomao/gin-budget-control/docs"
	"github.com/RaphaSalomao/gin-budget-control/model"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	SrvPort string
)

func RunServer() {
	router := gin.Default()

	groupUnauthorizedHandlers(router)
	groupReceiptHandlers(router)
	groupExpenseHandlers(router)
	groupSummaryHandlers(router)

	go router.Run(SrvPort)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Server is running")
	<-quit
}

func groupSummaryHandlers(router *gin.Engine) {
	summary := router.Group("/budget-control/api/v1/summary", authMiddleware())
	{
		summary.GET("/:year/:month", controller.MonthBalanceSumary)
	}
}

func groupExpenseHandlers(router *gin.Engine) {
	expense := router.Group("/budget-control/api/v1/expense", authMiddleware())
	{
		expense.POST("", validatorMiddleware[*model.ExpenseRequest](), controller.CreateExpense)
		expense.GET("", controller.FindAllExpenses)
		expense.GET("/:id", controller.FindExpense)
		expense.PUT("/:id", validatorMiddleware[*model.ExpenseRequest](), controller.UpdateExpense)
		expense.DELETE("/:id", controller.DeleteExpense)
		expense.GET("/period/:year/:month", controller.ExpensesByPeriod)
	}
}

func groupReceiptHandlers(router *gin.Engine) {
	receipt := router.Group("/budget-control/api/v1/receipt", authMiddleware())
	{
		receipt.POST("", validatorMiddleware[*model.ReceiptRequest](), controller.CreateReceipt)
		receipt.GET("", controller.FindAllReceipts)
		receipt.GET("/:id", controller.FindReceipt)
		receipt.PUT("/:id", validatorMiddleware[*model.ReceiptRequest](), controller.UpdateReceipt)
		receipt.DELETE("/:id", controller.DeleteReceipt)
		receipt.GET("/period/:year/:month", controller.ReceiptsByPeriod)
	}
}

func groupUnauthorizedHandlers(router *gin.Engine) {
	unauthorized := router.Group("/budget-control/api/v1")
	{
		unauthorized.POST("/user", validatorMiddleware[*model.UserRequest](), controller.CreateUser)
		unauthorized.POST("/user/authenticate", validatorMiddleware[*model.UserRequest](), controller.Authenticate)
		unauthorized.GET("/health", controller.Health)
		unauthorized.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}
