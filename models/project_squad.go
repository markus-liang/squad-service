package models

import (
	"time"
)

/* ATTRIBUTES */
// PUBLIC

//Project model definition
type ProjectSquad struct {
	ID          uint       `json:"-" gorm:"primaryKey"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`	
	ProjectID    uint    `json:"-" gorm:"not null"`
	Email        string  `json:"email" gorm:"size:50;not null;"`
	Status       string  `json:"status" gorm:"size:1;"`
	ChipinAmount int     `json:"chipin_amount"`
	Project      Project `json:"-"`
}
