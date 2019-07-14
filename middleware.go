package main

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
			c.Set("ticket", ticketInfo)
			c.Next()
		} else {
			c.JSON(200, gin.H{
				"result": "fake",
				"info":   "POST params wrong",
			})
		}
	}
}

// IsInDBMiddleware ...
func IsInDBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ticketInfo := c.MustGet("ticket").(TicketInfo)
		db := c.MustGet("DB").(*gorm.DB)
		tickets := Tickets{}
		value := db.Where(&Tickets{Key: ticketInfo.Key}).First(&tickets)
		if value.Error != nil {
			c.JSON(200, gin.H{
				"result":    "fake",
				"info":      "IsInDBMiddleware: DB Query error",
				"errorInfo": value.Error,
			})
		} else {
			if tickets.Times >= 0 {
				c.Set("ticketModel", tickets)
				c.Next()
			} else {
				c.JSON(200, gin.H{
					"result": "fake",
					"info":   "IsInDBMiddleware: error",
				})
			}
		}
	}
}

// ResultMiddleware ...
func ResultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.MustGet("resultCode").(string)
		c.Next()
	}
}
