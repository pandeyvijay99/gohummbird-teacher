package models

import "github.com/jinzhu/gorm"

// User struct
type Teacher struct {
	gorm.Model
	UserId string `gorm:"unique_index;not null" json:"user_id"`
	Name    string `gorm:"unique_index;not null" json:"name"`
	Designation string `gorm:"not null" json:"degination"`
	Image string `gorm:"not null" json:"image"`
	Description string `gorm:"not null" json:"description"`
	Qualification string `gorm:"not null" json:"qualification"`
	Facebook string `gorm:"not null" json:"facebook"`
	Google string `gorm:"not null" json:"google"`
	Twitter string `gorm:"not null" json:"twitter"`
	CreatedBy *time.Date `gorm:"not null" json:"created_by"`
}
