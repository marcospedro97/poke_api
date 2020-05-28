package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"poke/models"
)

func PokemonsIndex(c *gin.Context) {
	p := models.Pokemons{}
	models.All(&p)
	fmt.Println(p)
	j, err := json.Marshal(p)
	if err != nil {
		panic("erro no marshal de index")
	}
	c.String(http.StatusOK, string(j))
}

func PokemonsCreate(c *gin.Context) {

}
