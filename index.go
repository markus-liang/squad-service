package main

import (
	"log"
	auth "squad-service/controllers/auth"
	db "squad-service/databases"
	mid "squad-service/middlewares"
	mod "squad-service/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	con, err := mod.InitDB()
	defer con.Close()
	if err != nil {
		log.Panic(err)
	}

	db.Migrate(con)

	r := gin.Default()
	r.Use(mid.DBMiddleware(con))
	r.Use(mid.AuthMiddleware())

	r.POST("/login", auth.Login)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
