package models

import (
	"time"
)

/* ATTRIBUTES */
// PUBLIC

//Project model definition
type Project struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`	
	UserID      uint 	   `json:"-" gorm:"not null"`
	Leader		string 	   `json:"squad_leader" gorm:"size:50;"`
	Name        string     `json:"name" gorm:"size:50;not null;"`
	Description string     `json:"description" gorm:"size:1000;"`
	EndDate     time.Time  `json:"end_date" gorm:"not null;"`
	Amount      int        `json:"amount" gorm:"not null;"`
	User        User       `json:"-" gorm:"foreignKey:UserID"`
	ProjectSquad		[]ProjectSquad `json:"squad_member" gorm:"foreignkey:ProjectID"`
}
