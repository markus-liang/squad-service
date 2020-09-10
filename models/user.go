package models

import (
	"github.com/jinzhu/gorm"
)

/* ATTRIBUTES */
// PUBLIC

//User model definition
type User struct {
	gorm.Model
	Email    string `gorm:"size:50;unique_index"`
	Name     string `gorm:"size:50"`
	Password string `gorm:"size:100"`
	Status   string `gorm:"size:1"`
}

// UserStatus posibility values : (A)ctive, (I)nactive
var UserStatus = &userStatusList{
	Active:   "A",
	Inactive: "I",
}

// PRIVATE
type userStatusList struct {
	Active   string
	Inactive string
}
