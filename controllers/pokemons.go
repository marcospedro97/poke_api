package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"poke/models"
	"strconv"
)

func PokemonsIndex(c *gin.Context) {
	p := models.Pokemons{}
	json, err := models.All(&p)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.String(200, json)
}

func PokemonShow(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(406, gin.H{"error": err.Error()})
	}
	pokemon := models.Pokemon{}
	json, err := pokemon.Find(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
	}
	c.String(http.StatusOK, json)
}

func PokemonCreate(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	uuidString := fmt.Sprintf("%v", c.MustGet("userUUID"))
	pokemon := models.Pokemon{}
	pokemon.UserUUID, err = strconv.ParseUint(uuidString, 10, 32)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &pokemon); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	pokemon.Create()
	c.JSON(201, string(body))
}
