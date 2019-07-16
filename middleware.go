package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// SafeFilterMiddleware ...
func SafeFilterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ticketInfo TicketInfo
		c.BindJSON(&ticketInfo) // Bind json just form HTTP Post.
		keyStatus, _ := CheckKey(ticketInfo.Key)
		tokenStatus, _ := CheckToken(ticketInfo.Token)
		if keyStatus && tokenStatus {
			c.Set("ticket", ticketInfo)
			c.Next()
		} else {
			CreateLog(c.MustGet("DB").(*gorm.DB), ticketInfo.Key, 2, "Fake: SafeFilterMiddleware: Non-conformity")
			c.JSON(200, gin.H{
				"result": "fake",
				"info":   "SafeFilterMiddleware: Non-conformity",
			})
			c.Abort()
		}
	}
}

// SafeFilterMiddlewareByGet ...
func SafeFilterMiddlewareByGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ticketInfo TicketInfo
		ticketInfo.Key = c.Param("key")
		ticketInfo.Token = "4D96ADBB-C164-4AC4-A6FA-4513A0272F05"
		keyStatus, _ := CheckKey(ticketInfo.Key)
		tokenStatus, _ := CheckToken(ticketInfo.Token)
		if keyStatus && tokenStatus {
			c.Set("ticket", ticketInfo)
			c.Next()
		} else {
			CreateLog(c.MustGet("DB").(*gorm.DB), ticketInfo.Key, 2, "Fake: SafeFilterMiddleware: Non-conformity")
			c.JSON(200, gin.H{
				"result": "fake",
				"info":   "SafeFilterMiddlewareByGet: Non-conformity",
			})
			c.Abort()
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
			CreateLog(c.MustGet("DB").(*gorm.DB), ticketInfo.Key, 2, "Fake: SafeIsInDBMiddleware: DB Query error")
			c.JSON(200, gin.H{
				"result":    "fake",
				"info":      "SafeIsInDBMiddleware: DB Query error",
				"errorInfo": value.Error,
			})
			c.Abort()
		} else {
			if tickets.Times >= 0 {
				c.Set("ticketModel", tickets)
				c.Next()
			} else {
				CreateLog(c.MustGet("DB").(*gorm.DB), ticketInfo.Key, 2, "Fake: SafeIsInDBMiddleware: tickets.Times NaN")
				c.JSON(200, gin.H{
					"result": "fake",
					"info":   "SafeIsInDBMiddleware: tickets.Times NaN",
				})
				c.Abort()
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
			CreateLog(c.MustGet("DB").(*gorm.DB), tickets.Key, 3, "fuckyou: SafeIsStaffMiddleware: is not Staff")
			c.JSON(200, gin.H{
				"result": "fuckyou",
				"info":   "SafeIsStaffMiddleware: is not Staff",
			})
			c.Abort()
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
			CreateLog(c.MustGet("DB").(*gorm.DB), tickets.Key, 3, "fuckyou: SafeIsTicketMiddleware is not Ticket")
			c.JSON(200, gin.H{
				"result": "fuckyou",
				"info":   "SafeIsTicketMiddleware is not Ticket",
			})
			c.Abort()
		}
	}
}

// SafeIsInvalidMiddleware ...
func SafeIsInvalidMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tickets := Tickets{}
		tickets = c.MustGet("ticketModel").(Tickets)
		if tickets.Times == 0 || tickets.Times < 0 {
			CreateLog(c.MustGet("DB").(*gorm.DB), tickets.Key, 1, "invalid: SafeIsInvalidMiddleware : invalid!")
			c.JSON(200, gin.H{
				"result": "invalid",
				"info":   "SafeIsInvalidMiddleware : invalid!",
			})
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// SafeCertPictureMiddleware ..
func SafeCertPictureMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("picture")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			c.Abort()
			return
		}
		filename := uuid.Must(uuid.NewV4()).String()
		if err := c.SaveUploadedFile(file, "assets/"+filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		c.JSON(200, gin.H{
			"f": file.Header.Get("Content-Type"),
		})
		c.Next()
	}
}
