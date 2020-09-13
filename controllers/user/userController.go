package user

import (
	"net/http"
	"strconv"
	"time"
	"fmt"
	"reflect"

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
		h.RespondWithError(c, http.StatusUnprocessableEntity, "Invalid json input provided")
		return
	}

	if db.Where("email = ?", userParam.Email).First(&userDB); (reflect.DeepEqual(userDB, m.User{})) {
		h.RespondWithError(c, http.StatusUnauthorized, "User not found")
		return
	} else if userDB.Password != userParam.Password {
		h.RespondWithError(c, http.StatusUnauthorized, "Wrong password")
		return
	} else if userDB.Status != m.UserStatus.Active {
		h.RespondWithError(c, http.StatusUnauthorized, "User is inactive, contact your administrator")
		return
	}

	token, err := createToken(userDB)
	if err != nil {
		h.RespondWithError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := saveAuth(userDB.ID, token, redis)
	if saveErr != nil {
		h.RespondWithError(c, http.StatusUnprocessableEntity, saveErr.Error())
		return
	}	

	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}

	h.RespondSuccess(c, tokens)
}

func TestAuth(c *gin.Context){
	// user_email := c.MustGet("user_email").(string)
	access_uuid := c.MustGet("access_uuid").(string)
	h.RespondSuccess(c, "access_uuid = " + access_uuid)
}

func Activate(c *gin.Context){
	h.RespondWithError(c, http.StatusUnprocessableEntity, "To be implemented")
}

func SetPassword(c *gin.Context){
	h.RespondWithError(c, http.StatusUnprocessableEntity, "To be implemented")
}

func Signout(c *gin.Context){
	redis := c.MustGet("redis").(*redis.Client)
	access_uuid := c.MustGet("access_uuid").(string)

	_, err := redis.Del(access_uuid).Result()
	if err != nil {
		h.RespondWithError(c, http.StatusUnprocessableEntity, err)
		return
	}
	h.RespondSuccess(c, nil)
}

func Signup(c *gin.Context){
	var user m.User
	if err := c.ShouldBindJSON(&user); err != nil {
		h.RespondWithError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := user.Validate(c, "signup"); err != nil {
		h.RespondWithError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	user.Status = m.UserStatus.Inactive
	db := c.MustGet("db_mysql").(*gorm.DB)
	if err := db.Create(&user); err.Error != nil{
		fmt.Println("save");
		fmt.Println(err);		
		h.RespondWithError(c, http.StatusUnprocessableEntity, err.Error)
		return
	}

	h.RespondSuccess(c, nil)

	//h.RespondWithError(c, http.StatusUnprocessableEntity, "To be implemented")
}

func RequestCode(c *gin.Context){
	h.RespondWithError(c, http.StatusUnprocessableEntity, "To be implemented")
}

func VerifyCode(c *gin.Context){
	h.RespondWithError(c, http.StatusUnprocessableEntity, "To be implemented")
}

/////////////
// PRIVATE //
/////////////

// saveAuth is a function to save created token to redis
func saveAuth(userid uint64, td *m.TokenDetail, redis *redis.Client) error {
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
	td.AtExpires = time.Now().Add(time.Minute * 150).Unix()
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
