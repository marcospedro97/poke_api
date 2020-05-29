package main

import (
	"poke/models"
)

func main() {
	models.Init()
	models.Db.C.CreateTable(&models.Pokemon{})
	bulba := models.Pokemon{Code: 1, Name: "Bulbasaur", Type: "Grass, poison", NextEv: 2, PreviousEv: 0}
	bulba.Create()
}
