package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health Check
// @Description return server status
// @Tags Health
// @Success 200
// @Failure 404
// @Router /budget-control/api/v1/health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"online": true,
	})
}
