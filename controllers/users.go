package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"poke/middlewares"
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
	err = user.Authenticate()
	if err != nil {
		c.String(http.StatusUnauthorized, "user not authenticated")
	}
	token, err := middlewares.GenerateToken(user.Uuid)
	if err != nil {
		c.String(http.StatusInternalServerError, "{error: token not created}")
	}
	c.String(http.StatusOK, token)
}
