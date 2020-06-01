package main

import (
	"poke/models"
)

func main() {
	models.Init()
	models.Db.C.CreateTable(&models.Pokemon{})
	models.Db.C.CreateTable(&models.User{})
}
