package main

import "github.com/gin-gonic/gin"

// SafeIsTicketMiddleware ...
func SafeIsTicketMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tickets := Tickets{}
		tickets = c.MustGet("ticketModel").(Tickets)
		if tickets.Type == 1 || tickets.Type == 2 {
			c.Next()
		} else {
			c.JSON(200, gin.H{})
		}
	}
}
