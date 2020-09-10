package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

/* ATTRIBUTES */
// PUBLIC

//Project model definition
type Project struct {
	gorm.Model
	UserID      uint
	Name        string    `json:"name" gorm:"size:50;not null;"`
	Desctiption string    `json:"description" gorm:"size:1000;"`
	EndDate     time.Time `json:"enddate" gorm:"not null;"`
	Amount      int       `json:"amount" gorm:"not null;"`
	User        User      `gorm:"foreignkey:UserID"`
}
