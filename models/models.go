package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type conn struct {
	C *gorm.DB
}

var Db conn

func Init() {
	var err error
	Db.C, err = gorm.Open("sqlite3", "./database/test.db")
	if err != nil {
		panic("failed to connect database")
	}
}
