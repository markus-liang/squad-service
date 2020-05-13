package middlewares

import "github.com/gin-gonic/gin"

////////////
// PUBLIC //
////////////

// AuthMiddleware to authorize every requests
func AuthMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	//isAuthrized, message := authorize()
	isAuhtorized := true
	message := ""
	return func(c *gin.Context) {
		if !isAuhtorized {
			respondWithError(c, 401, message)
			return
		}
		c.Next()
	}
}

/////////////
// PRIVATE //
/////////////

// Authorize middleware for requests authorization
func authorize() (bool, string) {
	// NO IMPLEMENTED YET!

	return true, "Not Implemented Yet!"
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
