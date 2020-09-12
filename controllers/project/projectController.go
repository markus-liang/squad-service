package project

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "squad-service/models"
	h "squad-service/helpers"
)

// PUBLIC 

// Create (Public) endpoint to create a new project
func Create(c *gin.Context){
	db := c.MustGet("db_mysql").(*gorm.DB)
	userID := c.MustGet("user_id").(uint64)
	userEmail := c.MustGet("user_email").(string)

	var body interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		h.RespondWithError(c, http.StatusUnprocessableEntity, err.Error)
		return
	}

	input := body.(map[string]interface{})
	endDate,_ := time.Parse(time.RFC3339, input["end_date"].(string))

	var squads []m.ProjectSquad
	for _, j := range input["squad_member"].([]interface{}){

		squads = append(squads, m.ProjectSquad{
			Email: j.(map[string]interface{})["email"].(string),
			Status: m.ProjectSquadStatus.Invited,			
		})
	}

	data := m.Project{
		UserID: userID,
		Leader: userEmail,
		Name: input["name"].(string),
		Description: input["description"].(string),
		EndDate: endDate,
		Amount: int(input["amount"].(float64)),
		ProjectSquads: squads,
	}
	
	if err := db.Create(&data); err.Error != nil{
		h.RespondWithError(c, http.StatusUnprocessableEntity, err.Error)		
	}

	h.RespondSuccess(c, nil)
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
