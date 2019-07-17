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
	tickets := c.MustGet("ticketModel").(Tickets)
	tickets.Times = tickets.Times - 1
	ctrl.server.db.Save(&tickets)
	CreateLog(c.MustGet("DB").(*gorm.DB), tickets.Key, 0, "Ticket verification!")
	c.JSON(200, gin.H{
		"result": "success",
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
	tickets := c.MustGet("ticketModel").(Tickets)
	var staffPictures []StaffPicture
	result := ctrl.server.db.Where(&StaffPicture{Key: tickets.Key}).Find(&staffPictures)

	if result.Error != nil {
		c.JSON(200, gin.H{
			"result": "faild",
		})
		c.Abort()
	}
	CreateLog(c.MustGet("DB").(*gorm.DB), tickets.Key, 0, "success : GET STAFF PICTURE")
	c.JSON(200, gin.H{
		"result":   "success",
		"quantity": result.RowsAffected,
		"times":    tickets.Times,
		"records":  staffPictures,
	})
}

// Post ...
func (ctrl *StaffController) Post(c *gin.Context) {
	tickets := c.MustGet("ticketModel").(Tickets)
	staffPicture := c.MustGet("staffPicture").(StaffPicture)

	ctrl.server.db.Create(&StaffPicture{
		Key:  tickets.Key,
		Path: staffPicture.Path,
	})
	CreateLog(c.MustGet("DB").(*gorm.DB), tickets.Key, 0, "success : UPLOAD STAFF PICTURE "+staffPicture.Path)
	c.JSON(200, gin.H{
		"result": "success",
	})
}
