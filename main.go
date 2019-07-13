package main

/**
Thank you for working with me and teaching me to write go in the right way. > <
**/

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Server ...
type Server struct {
	isDebug  bool
	db       *gorm.DB
	server   *gin.Engine
	dbString string
}

// TicketInfo ...
type TicketInfo struct {
	Key   string `json:"key" form:"key"`
	Token string `json:"token" form:"token"`
}

func (server *Server) createDB() {
	connectionString := os.Getenv("ETS_DB")
	if len(connectionString) == 0 {
		connectionString = "root:1234@tcp(127.0.0.1:3306)/hitcomic?charset=utf8&parseTime=True"
	}
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("[x]something wrong in createDB() of server init.")
		panic(err)
	}
	db.LogMode(server.isDebug)

	db.AutoMigrate(&Tickets{})
	db.AutoMigrate(&Logs{})
	db.AutoMigrate(&CertPicture{})

	server.db = db
}

func (server *Server) createServer() {
	myServer := gin.Default()
	server.server = myServer

	// Middleware
	myServer.Use(func(c *gin.Context) {
		c.Set("DB", server.db)
	})

	myServer.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	myServer.POST("/ticket", func(c *gin.Context) {
		var ticketInfo TicketInfo
		c.BindJSON(&ticketInfo)

		c.JSON(200, gin.H{
			"key":   ticketInfo.Key,
			"token": ticketInfo.Token,
		})
	})
	myServer.GET("/test", func(c *gin.Context) {
		service := VerifyService{}
		c.JSON(200, gin.H{
			"message": service.VerifyTicket(server.db, "test"),
		})
	})
	myServer.Static("/assets", "./assets")
}

func main() {
	fmt.Println("dbString example: root:my-secret-pw@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True")
	ticketServer := &Server{}
	ticketServer.isDebug = len(os.Getenv("DUAN_DEBUG")) > 0
	ticketServer.createDB()
	ticketServer.createServer()
	ticketServer.server.Run()
}
