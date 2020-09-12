package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"github.com/go-redis/redis/v7"
	m "squad-service/models"
	h "squad-service/helpers"
)

////////////
// PUBLIC //
////////////

// AuthMiddleware to authorize every requests
func AuthMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	return func(c *gin.Context) {
		isAuthorized, access, message := authorize(c)		

		if !isAuthorized{
			h.RespondWithError(c, 401, message)
			return
		}

		// add context
		c.Set("user_id", access.UserID)
		c.Set("user_email", access.UserEmail)
		c.Set("access_uuid", access.AccessUUID)
		c.Next()
	}
}

/////////////
// PRIVATE //
/////////////

// Authorize middleware for requests authorization
func authorize(c *gin.Context) (bool, *m.AccessDetails, string) {
	access, err := extractTokenMetadata(c)
	if err != nil {
		return false, nil, err.Error()
	}

	isExist, err := fetchAuthFromRedish(c, access)
	if err != nil {
		return false, nil, "user has logged out"
	}

	if c.Request.Header.Get("user_email") != access.UserEmail {
		return false, nil, "wrong user"
	}

	return isExist, access, "OK"
}

func extractToken(c *gin.Context) string {
	bearToken := c.Request.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func extractTokenMetadata(c *gin.Context) (*m.AccessDetails, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		email, ok := claims["email"].(string)
		if !ok {
			return nil, err
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &m.AccessDetails{
			AccessUUID: accessUUID,
			UserID: userID,
			UserEmail: email,
		}, nil
	}
	return nil, err
}

func fetchAuthFromRedish(c *gin.Context, authD *m.AccessDetails) (bool, error) {
	redis := c.MustGet("redis").(*redis.Client)

	_, err := redis.Get(authD.AccessUUID).Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

func verifyToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
		return []byte(h.Env("JWT_ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
