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
	tickets := Tickets{}
	tickets = c.MustGet("ticketModel").(Tickets)
	tickets.Times = tickets.Times - 1
	ctrl.server.db.Save(&tickets)
	CreateLog(c.MustGet("DB").(*gorm.DB), tickets.Key, 0, "success")
	c.JSON(200, gin.H{
		"message": "success",
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

// Get ...
func (ctrl *StaffController) Get(c *gin.Context) {
	tickets := Tickets{}
	tickets = c.MustGet("ticketModel").(Tickets)
	var certPictures []CertPicture
	result := ctrl.server.db.Where(&CertPicture{Key: tickets.Key}).Find(&certPictures)

	if result.Error != nil {
		c.JSON(200, gin.H{
			"result":  "fake",
			"message": "DB query error",
		})
		c.Abort()
	}

	c.JSON(200, gin.H{
		"result":   "success",
		"quantity": result.RowsAffected,
		"times":    tickets.Times,
		"records":  certPictures,
	})
}

// Post ...
func (ctrl *StaffController) Post(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "it seems works",
	})
}
