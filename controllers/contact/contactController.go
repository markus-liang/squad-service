package contact

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	m "squad-service/models"
	h "squad-service/helpers"
)

// PUBLIC 

// List (Public) endpoint to get user's contact list
func List(c *gin.Context){
	db := c.MustGet("db_mysql").(*gorm.DB)
	user_id := c.MustGet("user_id").(uint64)

	var contacts []m.Contact
	if err := db.Where("deleted_at IS NULL AND user_id = ?", user_id).Find(&contacts); err.Error != nil {
		h.RespondWithError(c, http.StatusUnprocessableEntity, err.Error)
		return
	}

	h.RespondSuccess(c, contacts)
}