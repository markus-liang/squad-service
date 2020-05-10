package main

import (
	auth "tcf-service/services/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(authMiddleware())

	r.POST("/login", auth.Login)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func authMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	isAuthrized, message := auth.Authorize()

	return func(c *gin.Context) {
		if !isAuthrized {
			respondWithError(c, 401, message)
			return
		}
		c.Next()
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
