package main

import (
	"log"
	
	"squad-service/controllers/user"
	"squad-service/controllers/contact"
	"squad-service/controllers/project"
	db "squad-service/databases"
	h "squad-service/helpers"
	mid "squad-service/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// initiations
	con, err := initDB()
	defer con.Close()
	if err != nil {
		log.Panic(err)
	}

	redis, err := initRedis()
	if err != nil {
		log.Panic(err)
	}

	// migrations
	db.Migrate(con)

	r := gin.Default()

	// global middlewares
	r.Use(mid.CtxMiddleware(con, redis))

	// routes with no auth
/*
	r.GET("/users/activate", user.Activate)
	r.POST("/users/request_code", user.RequestCode)
*/
	r.POST("/users/signin", user.Signin)
	r.POST("/users/signup", user.Signup)
/*
	r.POST("/users/verify_code", user.VerifyCode)
*/
	// routes with auth 
	authGroup := r.Group("/")
	authGroup.Use(mid.AuthMiddleware())
	{
		authGroup.GET("/test", user.TestAuth)
		authGroup.GET("/contacts/list", contact.List)
		authGroup.GET("/projects/list", project.List)
		authGroup.POST("/projects/create", project.Create)
		authGroup.POST("/users/signout", user.Signout)
		/*
		authGroup.POST("/users/set_password", user.SetPassword)
		*/
	}


	r.Run() // listen and serve on 0.0.0.0:8080
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(h.Env("DB_DIALECT"), h.Env("DB_CONNECTION_STR"))
	if err != nil {
		return nil, err
	}
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func initRedis() (*redis.Client, error)  {
	//Initializing redis
	dsn := h.Env("REDIS_DSN")
	client := redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
