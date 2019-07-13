package main

import (
	"github.com/jinzhu/gorm"
	"regexp"
	"fmt"
)

// SafeService ...
type SafeService struct {
}
 
// QueryService ...
type QueryService struct{
	app *Server
}

// CheckService ...
type CheckService struct{
	app *Server
}

// VerifyService ...
type VerifyService struct{
}

// PictureService ...
type PictureService struct{
}

// CheckKey ...
func (service *SafeService) CheckKey(key string) (bool, error) {
	return regexp.MatchString("[a-zA-Z0-9]{32}", key)
}

// CheckToken ...
func (service *SafeService) CheckToken(token string) (bool, error) {
	return regexp.MatchString("\\w{8}(-\\w{4}){3}-\\w{12}", token)
}

// QueryTicket  ...
func (service *QueryService) QueryTicket(db *gorm.DB, key string) *gorm.DB{
	result := Tickets{Key: key}
	return db.First(&result) //Get a first matching record
}

// IsZero ...
func (service *CheckService) IsZero(db *gorm.DB, key string) bool{
	type Result struct{
		Key   string
		Type  uint
		Times uint
	}
	var result Result
	ticket := Tickets{Key: key}
	db.First(&ticket).Scan(&result) //Find Ticket times
	fmt.Println(result)
	if result.Times == 0 {
		return true
	}
	return false
}

//VerifyTicket ...
func (service *VerifyService) VerifyTicket(db *gorm.DB, key string) bool{
	type Result struct{
		Key   string
		Type  uint
		Times uint
	}
	var result Result
	ticket := Tickets{Key: key}
	db.First(&ticket).Scan(&result)
	db.Model(&ticket).First(&ticket).Update("Times", result.Times-1)
	return true
}

//VerifyStaff ...
func (service *VerifyService) VerifyStaff(db *gorm.DB, key string) bool{
	type Result struct{
		Key   string
		Type  uint
		Times uint
	}
	var result Result
	ticket := Tickets{Key: key}
	db.First(&ticket).Scan(&result)
	db.Model(&ticket).First(&ticket).Update("Times", result.Times-1)
	return true
}