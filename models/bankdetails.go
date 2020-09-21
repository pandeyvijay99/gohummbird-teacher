package models

import "github.com/jinzhu/gorm"
//import "time"
// User struct
type Bankdetails struct {
	gorm.Model
	
	UserId int 
	BankName string
	IfscCode string
	AccountNumber string
	UpiAddress string
	PanNumber string
	AdharNumber string
	NumberOfDependents int
	Religion string
	DateOfJoining string
	//Status string
	      
}



