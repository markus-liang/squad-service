package models

import "github.com/jinzhu/gorm"

// InitDB : initialize DB connection
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/tcf?parseTime=true")
	if err != nil {
		return nil, err
	}
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
