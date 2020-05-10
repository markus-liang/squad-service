package auth

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	"github.com/joho/godotenv"
)

var dbmap = initDb()

//////////////
// SERVICES //
//////////////

// Authorize (Public). A middleware for requests authorization
func Authorize() (bool, string) {
	// NO IMPLEMENTED YET!

	return false, "Not Implemented Yet!"
}

// Login (Public) endpoint to be called
func Login(c *gin.Context) {
	var userParam User
	var userDB User

	if err := c.ShouldBindJSON(&userParam); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json input provided")
		return
	}

	userDB, err := getUserByEmail(userParam.Email)

	if err == nil {
		if userDB.Password != userParam.Password {
			c.JSON(http.StatusUnauthorized, "Wrong password")
			return
		} else if userDB.Status != UserStatus.Active {
			c.JSON(http.StatusUnauthorized, "User is inactive, contact your administrator")
			return
		}

		token, err := createToken(userDB)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}

		c.JSON(http.StatusOK, token)
	} else {
		c.JSON(http.StatusFound, err.Error())
	}
}

/////////////
// ENGINES //
/////////////

func createToken(user User) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["email"] = user.Email
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(goDotEnvVariable("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

//getUserByEmail : function to get one user from DB
func getUserByEmail(email string) (User, error) {
	var userDB User

	err := dbmap.SelectOne(&userDB, "SELECT * FROM users WHERE email=?", email)

	return userDB, err
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/tcf?parseTime=true")
	checkErr(err, "sql.Open failed")

	dialect := gorp.MySQLDialect{"InnoDB", "UTF8"}
	dbmap := &gorp.DbMap{Db: db, Dialect: dialect}

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	return dbmap
}
