package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

////////////
// PUBLIC //
////////////

// DBMiddleware inject database connection as context
func DBMiddleware(con *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db_mysql", con)
		c.Next()
	}
}
