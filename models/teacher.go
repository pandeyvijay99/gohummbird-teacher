package models

import "github.com/jinzhu/gorm"
//import "time"
// User struct
type Teacher struct {
	gorm.Model
	//id	int `gorm:"unique_index;not null" json:"id"`
	UserId int 
	Name    string 
	Designation string 
	TeacherProfilePic string
	DateOfBirth string 
	Gender string
	Description string 
	PhoneNumber string
	PersonalEmail string
	Qualification string 
	Facebook string 
	Google string 
	Twitter string 
	Status  string          
	EmpID	string
	CreatedBy int        
	UpdatedBy int       
	DeletedBy int       
}


type Aser struct {
	gorm.Model
	id int
	UserId int
	Refer string
	Name  string
	DateOfBirth string
	Designation string
	teacherprofilepic string
	EmpID   string
	PersonalEmail string
	 PhoneNumber string 


}

type Arofile struct {
	gorm.Model
	id int
	Name      string
	User      Aser   `gorm:"association_foreignkey:Refer"` // use Refer as association foreign key
	UserRefer string `json:"user_refer"`
}
