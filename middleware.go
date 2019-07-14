package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SafeFilterMiddleware ...
func SafeFilterMiddleware() gin.HandlerFunc {
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
				"info":   "SafeFilterMiddleware: Non-conformity",
			})
		}
	}
}

// SafeIsInDBMiddleware ...
func SafeIsInDBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ticketInfo := c.MustGet("ticket").(TicketInfo)
		db := c.MustGet("DB").(*gorm.DB)
		tickets := Tickets{}
		value := db.Where(&Tickets{Key: ticketInfo.Key}).First(&tickets)
		if value.Error != nil {
			c.JSON(200, gin.H{
				"result":    "fake",
				"info":      "SafeIsInDBMiddleware: DB Query error",
				"errorInfo": value.Error,
			})
		} else {
			if tickets.Times >= 0 {
				c.Set("ticketModel", tickets)
				c.Next()
			} else {
				c.JSON(200, gin.H{
					"result": "fake",
					"info":   "SafeIsInDBMiddleware: tickets.Times NaN",
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

// SafeIsStaffMiddleware ...
func SafeIsStaffMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tickets := Tickets{}
		tickets = c.MustGet("ticketModel").(Tickets)
		if tickets.Type == 3 {
			c.Next()
		} else {
			c.JSON(200, gin.H{
				"result": "fuckyou",
				"info":   "SafeIsStaffMiddleware: is not Staff",
			})
		}
	}
}

// SafeIsTicketMiddleware ...
func SafeIsTicketMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tickets := Tickets{}
		tickets = c.MustGet("ticketModel").(Tickets)
		if tickets.Type == 1 || tickets.Type == 2 {
			c.Next()
		} else {
			c.JSON(200, gin.H{
				"result": "fuckyou",
				"info":   "SafeIsTicketMiddleware is not Ticket",
			})
		}
	}
}

// SafeIsInvalidMiddleware ...
func SafeIsInvalidMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tickets := Tickets{}
		tickets = c.MustGet("ticketModel").(Tickets)
		if tickets.Times == 0 || tickets.Times < 0 {
			c.JSON(200, gin.H{
				"result": "invalid",
				"info":   "SafeIsInvalidMiddleware : invalid!",
			})
		} else {
			c.Next()
		}
	}
}
