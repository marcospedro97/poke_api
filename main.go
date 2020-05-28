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
	r.Run()
}
