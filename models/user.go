package models

import (
	"errors"
	"strings"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

/* ATTRIBUTES */
// PUBLIC

//User model definition
type User struct {
	ID          uint64       `gorm:"primary_key" json:"-"`
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

func (u *User) Validate(c *gin.Context, action string) error{
	switch strings.ToLower(action) {
	case "signup":
		db := c.MustGet("db_mysql").(*gorm.DB)
		var result User

		if strings.TrimSpace(u.Name) == "" {
			return errors.New("Required name")
		}
		if len(strings.TrimSpace(u.Password)) < 8 {
			return errors.New("Required password with minimum 8 characters")
		}
		if strings.TrimSpace(u.Email) == "" {
			return errors.New("Required email")
		}
		if db.Where("email = ?", u.Email).First(&result); result.Email == u.Email {
			return errors.New("User already exists")
		}

		return nil
	default:
		return nil
	}

}


// PRIVATE
type userStatusList struct {
	Active   string
	Inactive string
}


