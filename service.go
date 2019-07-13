package main

import "regexp"

// SafeService ...
type SafeService struct {
}

// QueryService ...
type QueryService struct{
}

// CheckKey ...
func (service *SafeService) CheckKey(key string) (bool, error) {
	return regexp.MatchString("[a-zA-Z0-9]{32}", key)
}

// CheckToken ...
func (service *SafeService) CheckToken(token string) (bool, error) {
	return regexp.MatchString("\\w{8}(-\\w{4}){3}-\\w{12}", token)
}

// QueryTicket ...
func (service *QueryService) QueryTicket(key string) {
	var ticket []Tickets
	query := db.where("Key = ?",key).find(&ticket) //Get all matching records
	query.Error != nil{
		panic(query.Error)
	}
	return ticket
}
