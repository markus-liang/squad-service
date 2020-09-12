package models

import (
	"time"
)

/* ATTRIBUTES */
// PUBLIC

//Project model definition
type Project struct {
	ID          uint64     `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time  `json:"-" gorm:"default: NOW();"`
	UpdatedAt   time.Time  `json:"-" gorm:"default: NOW();"`
	DeletedAt   *time.Time `json:"-"`	
	UserID      uint64 	   `json:"-" gorm:"not null"`
	Leader		string 	   `json:"squad_leader" gorm:"size:50;"`
	Name        string     `json:"name" gorm:"size:50;not null;"`
	Description string     `json:"description" gorm:"size:1000;"`
	EndDate     time.Time  `json:"end_date" gorm:"not null;"`
	Amount      int        `json:"amount" gorm:"not null;"`
	User        User       `json:"-" gorm:"foreignKey:UserID"`
	ProjectSquads		[]ProjectSquad `json:"squad_member" gorm:"foreignkey:ProjectID"`
}
