package models

import (
	"time"
)

/* ATTRIBUTES */
// PUBLIC

//Project model definition
type ProjectSquad struct {
	ID           uint64     `json:"-" gorm:"primaryKey"`
	CreatedAt    time.Time  `json:"-"`
	UpdatedAt    time.Time  `json:"-"`
	DeletedAt    *time.Time `json:"-"`	
	ProjectID    uint64     `json:"-" gorm:"not null"`
	Email        string     `json:"email" gorm:"size:50;not null;"`
	Status       string     `json:"status" gorm:"size:1;"`
	ChipinAmount int        `json:"chipin_amount"`
	Project      Project    `json:"-"`
}

// ProjectStatus posibility values : (A)ctive, (I)nactive
var ProjectSquadStatus = &projectSquadStatusList{
	Invited: "I",
	Reject:	"R",
	WaitingPayment: "W",
	PaymentConfirmed: "C",
}

// PRIVATE
type projectSquadStatusList struct {
	Invited string
	Reject string
	WaitingPayment string
	PaymentConfirmed string
}
