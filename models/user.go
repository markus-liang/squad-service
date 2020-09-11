package models

import (
	"time"
)

/* ATTRIBUTES */
// PUBLIC

//User model definition
type User struct {
	ID          uint       `gorm:"primary_key" json:"-"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`	
	Email    string `gorm:"size:50;uniqueIndex"`
	Name     string `gorm:"size:50"`
	Password string `gorm:"size:100"`
	Status   string `gorm:"size:1"`
	Project	 []Project
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
