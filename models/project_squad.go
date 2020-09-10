package models

import (
	"github.com/jinzhu/gorm"
)

/* ATTRIBUTES */
// PUBLIC

//Project model definition
type ProjectSquad struct {
	gorm.Model
	ProjectID    uint
	Email        string  `json:"email" gorm:"size:50;not null;"`
	Status       string  `json:"status" gorm:"size:1;"`
	ChipinAmount int     `json:"chipin_amount"`
	Project      Project `gorm:"foreignkey:ProjectID"`
}
