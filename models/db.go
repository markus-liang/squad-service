package models

import (
	h "squad-service/helpers"

	"github.com/jinzhu/gorm"
)

// InitDB : initialize DB connection
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(h.Env("DB_DIALECT"), h.Env("DB_CONNECTION_STR"))
	if err != nil {
		return nil, err
	}
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
