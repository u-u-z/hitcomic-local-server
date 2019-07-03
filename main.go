package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server struct {
	isDebug  bool
	db       *gorm.DB
	server   *gin.Engine
	dbString string
}

func (server *Server) createDB() {
	connectionString := os.Getenv("ETS_DB")
	if len(connectionString) == 0 {
		connectionString = "root:my-secret-pw@tcp(127.0.0.1:3306)/hitcomic?charset=utf8&parseTime=True"
	}
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	db.LogMode(server.isDebug)

	// db.AutoMigrate(&URL{})

	server.db = db
}

func main() {
	fmt.Println("dbString example: root:my-secret-pw@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True")

	myServer := gin.Default()
	myServer.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	myServer.Run()
}
