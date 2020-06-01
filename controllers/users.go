package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"poke/models"
)

func UserCreate(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		panic("error in create User")
	}
	user := models.User{}
	if err := json.Unmarshal(body, &user); err != nil {
		c.String(http.StatusInternalServerError, "{error: 'invalid json'}")
	}
	user.Create()
}

func UserRead(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		panic("error in read User")
	}
	user := models.User{}
	if err = json.Unmarshal(body, &user); err != nil {
		c.String(http.StatusInternalServerError, "{error: 'invalid json'}")
	}
	if err = user.Authenticate(); err == nil {
		t := models.NewToken(user.Uuid)
		var jwt string
		jwt, err = t.GenerateJWT()
		c.String(http.StatusOK, jwt)
	}
}
