package main

import "regexp"

// SafeService ...
type SafeService struct {
}
 
// QueryService ...
type QueryService struct {
	app *Server
}

// CheckService ...
type CheckService struct{
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
func (service *QueryService) QueryTicket(key string) {
<<<<<<< HEAD

=======
	return Server.db.where("Key = ?",key).first(&ticket) //Get a first matching record
>>>>>>> 9ea79d3e1fd332b075988ea1f86bbf5d5beaa24c
}
