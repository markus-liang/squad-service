package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/go-redis/redis/v7"
)

////////////
// PUBLIC //
////////////

// CtxMiddleware inject needed contexts such as : database connection, redis
func CtxMiddleware(con *gorm.DB, redis *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db_mysql", con)
		c.Set("redis", redis)		
		c.Next()
	}
}
