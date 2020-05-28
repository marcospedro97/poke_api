package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Pokemon struct {
	gorm.Model `json:"-"`
	Code       int    `gorm:"unique" json:"code"`
	Name       string `gorm:"unique" json:"name"`
	Type       string `json:"type"`
	NextEv     int    `json:"Next_Evolution"`
	PreviousEv int    `json:"Previous_Evolution"`
}

type Pokemons []Pokemon

func (p *Pokemon) Create() {
	db.DB.Create(&p)
}

func All(p *Pokemons) {
	db.DB.Find(&p)
	fmt.Println(p)
}
