package main

import (
	"regexp"
	"time"

	"github.com/jinzhu/gorm"
)

// CheckKey ...
func CheckKey(key string) (bool, error) {
	return regexp.MatchString("[a-zA-Z0-9]{32}", key)
}

// CheckToken ...
func CheckToken(token string) (bool, error) {
	return regexp.MatchString("\\w{8}(-\\w{4}){3}-\\w{12}", token)
}

// CreateLog ...
func CreateLog(db *gorm.DB, key string, result uint, info string) {
	log := Logs{
		Key:    key,
		Result: result,
		Info:   info,
		BasicModel: BasicModel{
			CreatedAt: time.Now(),
		},
	}
	db.Create(&log)
}

// FindStaffPictureRowByTicketKey ...
func FindStaffPictureRowByTicketKey(db *gorm.DB, key string) {

}
