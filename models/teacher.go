package models

import "github.com/jinzhu/gorm"

// User struct
type Teacher struct {
	gorm.Model
	//id	int `gorm:"unique_index;not null" json:"id"`
	UserId int `gorm:"unique_index;not null" json:"user_id"`
	Name    string `gorm:"unique_index;not null" json:"name"`
	Designation string `gorm:"not null" json:"degination"`       
	teacherprofilepic string `gorm:"not null" json:"teacher_profile_pic"`
	Description string `gorm:"not null" json:"description"`
	Qualification string `gorm:"not null" json:"qualification"`
	PhoneNumber string `gorm:"not null" json:"phone_number"`
	Facebook string `gorm:"not null" json:"facebook"`
	Google string `gorm:"not null" json:"google"`
	Twitter string `gorm:"not null" json:"twitter"`
	Status           int8   `json:"status"`
	EmpID	string	`gorm:"not null" json:"emp_id"`
	CreatedBy        int8   `json:"created_by"`
	UpdatedBy        int8   `json:"updated_by"`
	DeletedBy        int8   `json:"deleted_by"`
}


type Aser struct {
	gorm.Model
	id int
	Refer string
	Name  string
}

type Arofile struct {
	gorm.Model
	id int
	Name      string
	User      Aser   `gorm:"association_foreignkey:Refer"` // use Refer as association foreign key
	UserRefer string `json:"user_refer"`
}
