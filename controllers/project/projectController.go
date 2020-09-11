package project

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	m "squad-service/models"
	h "squad-service/helpers"
)

// PUBLIC 

// Create (Public) endpoint to create a new project
func Create(c *gin.Context){
	db := c.MustGet("db_mysql").(*gorm.DB)
	user_id := c.MustGet("user_id").(uint64)

	var projects []m.Project
	if err := db.Preload("ProjectSquad").Where("user_id = ?", user_id).Find(&projects); err.Error != nil {
		h.RespondWithError(c, http.StatusUnprocessableEntity, err.Error)
		return
	}

	h.RespondSuccess(c, projects)
}

// List (Public) endpoint to get user's project list
func List(c *gin.Context){
	db := c.MustGet("db_mysql").(*gorm.DB)
	user_id := c.MustGet("user_id").(uint64)

	var projects []m.Project
	if err := db.Preload("ProjectSquad").Where("user_id = ?", user_id).Find(&projects); err.Error != nil {
		h.RespondWithError(c, http.StatusUnprocessableEntity, err.Error)
		return
	}

	h.RespondSuccess(c, projects)
}
