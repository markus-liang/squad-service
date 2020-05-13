package models

import (
	"github.com/jinzhu/gorm"
)

/* ATTRIBUTES */
// PUBLIC

//User model definition
type User struct {
	gorm.Model
	Email    string `gorm:"size:50;unique_index" `
	Password string `gorm:"size:100"`
	Status   string `gorm:"size:1"`
	RoleID   uint
	Role     Role `gorm:"foreignkey:RoleID"`
}

// UserStatus posibility values : (A)ctive, (I)nactive, (B)anned
var UserStatus = &userStatusList{
	Active:   "A",
	Inactive: "I",
	Banned:   "B",
}

// PRIVATE
type userStatusList struct {
	Active   string
	Inactive string
	Banned   string
}
