package main

import (
	"github.com/gin-gonic/gin"
)

// TicketController ...
type TicketController struct {
	server *Server
}

// Get ...
func (ctrl *TicketController) Get(c *gin.Context) {

}

// Post ...
func (ctrl *TicketController) Post(c *gin.Context) {
	ticketInfo := c.MustGet("ticket").(TicketInfo)
	//ctrl.TicketInfo.Key = ticketInfo.Key
	c.JSON(200, gin.H{
		"message": ctrl.server.db.First(&Tickets{Key: ticketInfo.Key}),
	})
}
