package models

import "github.com/jinzhu/gorm"
//import "time"
// User struct
type TeacherTimeTable struct {
	gorm.Model
	
	UserId int 
	ClassId int
	SectionId int
	SubjectId int
	IsPrimary string
	
	//Status string
	      
}



