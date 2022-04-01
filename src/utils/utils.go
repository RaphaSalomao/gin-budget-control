package utils

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/RaphaSalomao/gin-budget-control/model"
	"github.com/gin-gonic/gin"
)

func MonthIntervalFrom(date string) (firstDay, lastDay time.Time, err error) {
	year, month, err := GetYearMonthFromDateString(date)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	firstDay = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	lastDay = time.Date(year, time.Month(month+1), 1, 0, 0, 0, -1, time.Local)
	return
}

func GetYearMonthFromDateString(date string) (int, int, error) {
	splitDate := strings.Split(date, "-")
	year, err := strconv.Atoi(splitDate[0])
	if err != nil {
		return -1, -1, errors.New("unable to parse date")
	}
	month, err := strconv.Atoi(splitDate[1])
	if err != nil {
		return -1, -1, errors.New("unable to parse date")
	}
	return year, month, nil
}

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}

func IsHashAndPasswordMatched(password string, hash string) bool {
	return HashPassword(password) == hash
}

func RespondWithError(c *gin.Context, status int, message string, shouldAbort bool, body ...interface{}) {
	c.JSON(status, model.ErrorResponse{
		Error: message,
		Body:  body,
	})
	if shouldAbort {
		c.Abort()
	}
}
