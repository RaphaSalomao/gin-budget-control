package router

import (
	"net/http"
	"strings"

	"github.com/RaphaSalomao/gin-budget-control/model"
	"github.com/RaphaSalomao/gin-budget-control/utils"
	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Split(c.Request.Header.Get("Authorization"), " ")
		if len(token) != 2 {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Error:   "Unauthorized",
				Message: "Invalid token",
			})
			c.Abort()
			return
		}
		if token[0] != "Bearer" || len(token) != 2 {
			c.AbortWithStatus(400)
			return
		}
		userId, err := utils.ParseToken(token[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Error:   "Unauthorized",
				Message: "Invalid token",
			})
			c.Abort()
			return
		}
		c.Set("userId", userId)
		c.Next()
	}
}
