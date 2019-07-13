package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	tickets := Tickets{}
	value := ctrl.server.db.Where(&Tickets{Key: ticketInfo.Key}).First(&tickets).Value.(*Tickets)
	fmt.Println(value)
	fmt.Println(tickets)
	if value.Times > 1 {
		c.JSON(200, gin.H{
			"message": value.Times,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "fuck wrong",
		})
	}
	/*
		c.JSON(200, gin.H{
			"message": ctrl.server.db.First(&Tickets{Key: ticketInfo.Key}).Value,
		})*/
}

// GetTicketTimes ...
func (ctrl *TicketController) GetTicketTimes(db *gorm.DB, key string) int {
	return 0
}
