package router

import (
	"net/http"
	"strings"

	"github.com/RaphaSalomao/gin-budget-control/model"
	"github.com/RaphaSalomao/gin-budget-control/model/implement"
	"github.com/RaphaSalomao/gin-budget-control/security"
	"github.com/RaphaSalomao/gin-budget-control/utils"
	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")
		if len(token) != 2 {
			utils.RespondWithError(c, http.StatusUnauthorized, "Invalid Token", true)
			return
		}
		user, err := security.GetLoggedUserFromToken(token[1])
		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, "Invalid Token", true)
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func validatorMiddleware[T implement.Validable]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var validable T
		if err := c.ShouldBindJSON(&validable); err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid Request Body", true)
			return
		}
		if err := validable.Validate(); err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, err.Error(), true,
				model.ErrorResponse{
					Error:   "Validation Error",
					Message: err.Error(),
				},
			)
			return
		}
		c.Set("body", validable)
	}
}
