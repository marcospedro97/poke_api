package main

import (
	"github.com/gin-gonic/gin"
	"poke/controllers"
	"poke/models"
)

func main() {
	models.Init()
	r := gin.Default()
	r.GET("/", controllers.PokemonsIndex)
	r.GET("/:id", controllers.PokemonShow)
	r.POST("/create", controllers.PokemonCreate)
	r.POST("/sign_up", controllers.UserCreate)
	r.POST("/sign_in", controllers.UserRead)
	r.Run()
}
