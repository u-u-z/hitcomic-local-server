package main

import (
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
	c.JSON(200, gin.H{
		"message": "it works!",
	})
}

// GetTicketTimes ...
func (ctrl *TicketController) GetTicketTimes(db *gorm.DB, key string) int {
	return 0
}

// StaffController ...
type StaffController struct {
	server *Server
}

// Post ...
func (ctrl *StaffController) Post(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "fuck wrong",
	})
}
