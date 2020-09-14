package models

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-redis/redis/v7"
)

/* ATTRIBUTES */
// PUBLIC

//User model definition
type User struct {
	ID          uint64     `gorm:"primary_key" json:"-"`
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

func (u *User) GenerateActivationCode(c *gin.Context) (string, error) {
	redis := c.MustGet("redis").(*redis.Client)
	now := time.Now()

	rand.Seed(now.UnixNano())
	code := strconv.Itoa(rand.Intn(89999) + 10000)

	exp,_ := time.ParseDuration("24h")
    err := redis.Set("uact_" + u.Email + "_" + code, code, exp).Err()
    if err != nil {
        return "", err
    }
	return code, nil
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.Email = strings.TrimSpace(u.Email)
	u.Name = strings.TrimSpace(u.Name)
	u.Password = strings.TrimSpace(u.Password)

	if u.Status == "" {
		u.Status = UserStatus.Inactive
	}
}

func (u *User) Validate(c *gin.Context, action string) error{
	switch strings.ToLower(action) {
	case "signup":
		db := c.MustGet("db_mysql").(*gorm.DB)
		var result User

		if u.Name == "" {
			return errors.New("Required name")
		}
		if len(u.Password) < 8 {
			return errors.New("Required password with minimum 8 characters")
		}
		if u.Email == "" {
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

func (u *User) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// PRIVATE
type userStatusList struct {
	Active   string
	Inactive string
}
