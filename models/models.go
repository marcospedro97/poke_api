package models

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type conn struct {
	C *gorm.DB
	R *redis.Client
}

var Db conn

func Init() {
	var err error
	Db.C, err = gorm.Open("sqlite3", "./database/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	err = redisConnect()
	if err != nil {
		fmt.Println("failed on redis")
		panic("failed to connect redis")
	}
}

func redisConnect() error {
	Db.R = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := Db.R.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
