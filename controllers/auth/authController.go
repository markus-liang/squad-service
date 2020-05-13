package auth

import (
	"net/http"
	"time"

	c "tcf-service/controllers"
	m "tcf-service/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

////////////
// PUBLIC //
////////////

// Login (Public) endpoint to be called
func Login(c *gin.Context) {
	db := c.MustGet("db_mysql").(*gorm.DB)

	var userParam m.User
	var userDB m.User

	if err := c.ShouldBindJSON(&userParam); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json input provided")
		return
	}

	if db.Where("email = ?", userParam.Email).First(&userDB); (userDB == m.User{}) {
		c.JSON(http.StatusUnauthorized, "User not found")
		return
	} else if userDB.Password != userParam.Password {
		c.JSON(http.StatusUnauthorized, "Wrong password")
		return
	} else if userDB.Status != m.UserStatus.Active {
		c.JSON(http.StatusUnauthorized, "User is inactive, contact your administrator")
		return
	}

	token, err := createToken(userDB)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, token)
}

/////////////
// PRIVATE //
/////////////

func createToken(user m.User) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["email"] = user.Email
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(c.Env("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
