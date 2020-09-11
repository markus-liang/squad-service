package models

import (
	"time"
)

/* ATTRIBUTES */
// PUBLIC

//Contact model definition
type Contact struct {
	ID          uint       `gorm:"primary_key" json:"-"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`	
	UserID uint   `json:"-" gorm:"not null"`
	Email  string `json:"email" gorm:"size:50;not null;unique_index"`
	Alias  string `json:"alias" gorm:"size:50;not null"`
	Status string `json:"status" gorm:"size:1"`
	User   User   `json:"-" gorm:"foreignkey:UserID"`
}

// ContactStatus posibility values : (A)ctive, (I)nvited, (D)eleted
var ContactStatus = &contactStatusList{
	Active:   "A",
	Invited:  "I",
	Deleted:  "D",
}

// PRIVATE
type contactStatusList struct {
	Active  string
	Invited string
	Deleted string
}
