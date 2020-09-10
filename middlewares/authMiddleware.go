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
			respondWithError(c, 401, message)
			return
		}

		c.Set("user_email", access.UserEmail)
		c.Set("access_uuid", access.AccessUuid)
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
		return false, nil, "unauthorized : " + err.Error()
	}

	isExist, err := fetchAuthFromRedish(c, access)
	if err != nil {
		return false, nil, "unauthorized : user has logged out"
	}

	if c.Request.Header.Get("user_email") != access.UserEmail {
		return false, nil, "unauthorized : wrong user"
	}

	return isExist, access, "ok"
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
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		email, ok := claims["email"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &m.AccessDetails{
			AccessUuid: accessUuid,
			UserId: userId,
			UserEmail: email,
		}, nil
	}
	return nil, err
}

func fetchAuthFromRedish(c *gin.Context, authD *m.AccessDetails) (bool, error) {
	redis := c.MustGet("redis").(*redis.Client)

	_, err := redis.Get(authD.AccessUuid).Result()
	if err != nil {
		return false, err
	}
	//userID, _ := strconv.ParseUint(userid, 10, 64)
	return true, nil
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
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
