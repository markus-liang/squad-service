package user

import (
	"net/http"
	"strconv"
	"time"

	h "squad-service/helpers"
	m "squad-service/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"github.com/twinj/uuid"
)

////////////
// PUBLIC //
////////////

// Signin (Public) endpoint to be called
func Signin(c *gin.Context) {
	db := c.MustGet("db_mysql").(*gorm.DB)
	redis := c.MustGet("redis").(*redis.Client)

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

	saveErr := saveAuth(userDB.ID, token, redis)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}	

	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}

	c.JSON(http.StatusOK, tokens)
}

func TestAuth(c *gin.Context){
	// user_email := c.MustGet("user_email").(string)
	access_uuid := c.MustGet("access_uuid").(string)

	c.JSON(http.StatusOK, "access_uuid = " + access_uuid)
}

func Activate(c *gin.Context){
	c.JSON(http.StatusUnprocessableEntity, "To be implemented")
}

func SetPassword(c *gin.Context){
	c.JSON(http.StatusUnprocessableEntity, "To be implemented")
}

func Signout(c *gin.Context){
	redis := c.MustGet("redis").(*redis.Client)
	access_uuid := c.MustGet("access_uuid").(string)

	_, err := redis.Del(access_uuid).Result()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	c.JSON(http.StatusOK, "OK")

}

func Signup(c *gin.Context){
	c.JSON(http.StatusUnprocessableEntity, "To be implemented")
}

func RequestCode(c *gin.Context){
	c.JSON(http.StatusUnprocessableEntity, "To be implemented")
}

func VerifyCode(c *gin.Context){
	c.JSON(http.StatusUnprocessableEntity, "To be implemented")
}

/////////////
// PRIVATE //
/////////////

// saveAuth is a function to save created token to redis
func saveAuth(userid uint, td *m.TokenDetail, redis *redis.Client) error {
    at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
    rt := time.Unix(td.RtExpires, 0)
    now := time.Now()
 
    errAccess := redis.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
    if errAccess != nil {
        return errAccess
    }
    errRefresh := redis.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
    if errRefresh != nil {
        return errRefresh
    }
    return nil
}

// createToken is a function to generate new token
func createToken(user m.User) (*m.TokenDetail, error) {
	td := &m.TokenDetail{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	// creating access token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = user.ID
	atClaims["email"] = user.Email
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(h.Env("JWT_ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	// creating access token
	rtClaims := jwt.MapClaims{}
	rtClaims["authorized"] = true
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = user.ID
	rtClaims["email"] = user.Email
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.RefreshToken, err = rt.SignedString([]byte(h.Env("JWT_REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}
