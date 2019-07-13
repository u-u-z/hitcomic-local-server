package main

import (
	"regexp"

	"github.com/gin-gonic/gin"
)

// CheckKey ...
func CheckKey(key string) (bool, error) {
	return regexp.MatchString("[a-zA-Z0-9]{32}", key)
}

// CheckToken ...
func CheckToken(token string) (bool, error) {
	return regexp.MatchString("\\w{8}(-\\w{4}){3}-\\w{12}", token)
}

// SafeMiddleware ...
func SafeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ticketInfo TicketInfo
		c.BindJSON(&ticketInfo)
		keyStatus, _ := CheckKey(ticketInfo.Key)
		tokenStatus, _ := CheckToken(ticketInfo.Token)
		if keyStatus && tokenStatus {
			c.Next()
		} else {
			c.JSON(200, gin.H{
				"result": "fake",
				"info":   "POST params wrong",
			})
		}

	}
}
