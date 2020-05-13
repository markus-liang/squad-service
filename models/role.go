package models

import (
	"github.com/jinzhu/gorm"
)

/* ATTRIBUTES */
// PUBLIC

//Role model definition
type Role struct {
	gorm.Model
	Name string `json:"name" gorm:"size:20;not null;unique"`
}
