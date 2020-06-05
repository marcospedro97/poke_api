package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"poke/middlewares"
	"poke/models"
)

func UserCreate(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
	}
	user := models.User{}
	if err := json.Unmarshal(body, &user); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
	}
	user.Create()
	c.String(201, string(body))
}

func UserRead(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
	}
	user := models.User{}
	if err = json.Unmarshal(body, &user); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
	}
	err = user.Authenticate()
	if err != nil {
		c.JSON(401, gin.H{"errors": err.Error()})
	}
	token, err := middlewares.GenerateToken(user.Uuid)
	if err != nil {
		c.JSON(500, gin.H{"errors": err.Error()})
	}
	c.JSON(200, gin.H{"token": token})
}
