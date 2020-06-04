package middlewares

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	"net/http"
	"os"
	"poke/models"
	"strconv"
	"time"
)

type AccessDetails struct {
	AccessUUID string
	UserUUID   uint32
	Expires    int64
	jwt.StandardClaims
}

func GenerateToken(u uint32) (string, error) {
	var err error
	exp := time.Now().Add(time.Hour * 6)
	session := uuid.NewV4().String()
	os.Setenv("API_KEY", "ASDLKFAKSLDFKAJSHDFAJSHDLJFHASLKDFJHAL")
	claims := AccessDetails{UserUUID: u, AccessUUID: session, Expires: exp.Unix()}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rt.SignedString([]byte(os.Getenv("API_KEY")))
	if err != nil {
		return "", err
	}
	errorRedis := models.Db.R.Set(session, strconv.Itoa(int(u)), exp.Sub(time.Now())).Err()
	if errorRedis != nil {
		return "", errorRedis
	}
	return token, nil
}

func ValidateToken(c *gin.Context) {
	tknStr := c.GetHeader("Authentication")
	claims := &AccessDetails{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_KEY")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.String(http.StatusUnauthorized, "invalid signature")
			c.Abort()
		}
		c.String(http.StatusBadRequest, "invalid params")
		c.Abort()
	}
	if !tkn.Valid {
		c.String(http.StatusUnauthorized, "invalid token")
		c.Abort()
	}
	userID, err := models.Db.R.Get(claims.AccessUUID).Result()
	if err != nil {
		c.String(http.StatusUnauthorized, "session not found")
		c.Abort()
	}
	c.Set("userUUID", userID)
	c.Next()
}
