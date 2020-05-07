package main

import (
	"net/http"
	auth "tcf-service/services/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	/* sample routes */
	r.GET("/", index)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	/* end of sample routes */

	r.POST("/login", auth.Login)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "Hello brother!")
}
