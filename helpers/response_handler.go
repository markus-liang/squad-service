package helpers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mickep76/mapslice-json"
)

func RespondSuccess(c *gin.Context, data interface{}) {
	if data != nil{
		c.JSON(
			http.StatusOK, 
			mapslice.MapSlice{
				mapslice.MapItem{Key: "status", Value: "OK"},
				mapslice.MapItem{Key: "data", Value: data},
			},
		)		
	}else{
		c.JSON(
			http.StatusOK, 
			mapslice.MapSlice{
				mapslice.MapItem{Key: "status", Value: "OK"},
			},
		)
	}
}

func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(
		code, 
		mapslice.MapSlice{
			mapslice.MapItem{Key: "status", Value: "ERR"},
			mapslice.MapItem{Key: "msg", Value: message},
		},
	)
}
