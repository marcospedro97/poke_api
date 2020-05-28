package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type conn struct {
	DB *gorm.DB
}

var db conn

func Init() {
	var err error
	db.DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.DB.CreateTable(&Pokemon{})
	data := Pokemon{Code: 1, Name: "alph", Type: "fire", NextEv: 2, PreviousEv: 0}
	data.Create()

}
