package models

import (
	"github.com/jinzhu/gorm"
)

/* ATTRIBUTES */
// PUBLIC

//Contact model definition
type Contact struct {
	gorm.Model
	UserID uint
	Email  string `json:"email" gorm:"size:50;not null;unique_index"`
	Alias  string `json:"alias" gorm:"size:50;not null"`
	User   User   `gorm:"foreignkey:UserID"`
}
