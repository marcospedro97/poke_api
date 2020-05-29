package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"poke/models"
	"strconv"
)

func PokemonsIndex(c *gin.Context) {
	p := models.Pokemons{}
	json := models.All(&p)
	c.String(http.StatusOK, json)
}

func PokemonShow(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic("erro no parametro id")
	}
	pokemon := models.Pokemon{}
	json := pokemon.Find(id)
	c.String(http.StatusOK, json)
}

func PokemonCreate(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		panic("erro no create")
	}
	pokemon := models.Pokemon{}
	if err := json.Unmarshal(body, &pokemon); err != nil {
		panic("erro no unmarshal")
	}
	pokemon.Create()
	c.String(http.StatusOK, string(body))
}
