package models

import (
	"encoding/json"
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

func All(ps *Pokemons) string {
	Db.C.Find(&ps)
	json, err := json.Marshal(ps)
	if err != nil {
		panic("erro no list all")
	}
	return string(json)
}

func (p *Pokemon) Create() {
	Db.C.Create(&p)
}

func (p *Pokemon) Find(id int) string {
	Db.C.Where("Code = ?", id).Find(&p)
	json, err := json.Marshal(p)
	if err != nil {
		panic("erro no find")
	}
	return string(json)
}
