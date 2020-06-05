package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type Pokemon struct {
	gorm.Model `json:"-"`
	UserUUID   uint64 `json:"-"`
	Code       int    `gorm:"unique" json:"Code"`
	Name       string `gorm:"unique" json:"Name"`
	Type       string `json:"Type"`
	NextEv     int    `json:"Next_evolution"`
	PreviousEv int    `json:"Previous_evolution"`
}

type Pokemons []Pokemon

func All(ps *Pokemons) (string, error) {
	Db.C.Find(&ps)
	json, err := json.Marshal(ps)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func (p *Pokemon) Create() {
	Db.C.Create(&p)
}

func (p *Pokemon) Find(id int) (string, error) {
	Db.C.Where("Code = ?", id).Find(&p)
	json, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(json), nil
}
