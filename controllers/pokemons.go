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
	json := models.All(&p)
	c.String(http.StatusOK, json)
}

func PokemonShow(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.String(http.StatusNotFound, "{error: 'invalid id'}")
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
	uuidString := fmt.Sprintf("%v", c.MustGet("userUUID"))
	pokemon := models.Pokemon{}
	pokemon.UserUUID, err = strconv.ParseUint(uuidString, 10, 32)
	if err != nil {
		panic(err)
	}
	fmt.Println(pokemon.UserUUID)
	if err := json.Unmarshal(body, &pokemon); err != nil {
		c.String(http.StatusInternalServerError, "{error: 'invalid json'}")
	}
	pokemon.Create()
	c.String(http.StatusOK, string(body))
}
