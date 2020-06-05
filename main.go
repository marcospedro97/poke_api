package main

import (
	"github.com/gin-gonic/gin"
	"poke/controllers"
	"poke/middlewares"
	"poke/models"
)

func main() {
	models.Init()
	r := gin.Default()
	c := r.Group("/user")
	{
		r.GET("/", controllers.PokemonsIndex)
		r.GET("/:id", controllers.PokemonShow)
		c.Use(middlewares.ValidateToken)
		c.POST("/create", controllers.PokemonCreate)
	}
	r.POST("/sign_up", controllers.UserCreate)
	r.POST("/sign_in", controllers.UserRead)
	r.Run()
}
