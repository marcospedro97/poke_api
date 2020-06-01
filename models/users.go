package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model `json:"-"`
	Uuid       uint32 `json:"-" gorm:"primary_key;auto_increment:false"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"password" gorm:"not_null"`
}

func (u *User) Create() error {
	u.Uuid = uuid.New().ID()
	hash, err := hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	Db.C.Create(&u)
	return nil
}

func (u *User) Authenticate() error {
	t := User{}
	Db.C.Where("email = ?", u.Email).Find(&t)
	err := bcrypt.CompareHashAndPassword([]byte(t.Password), []byte(u.Password))
	if err != nil {
		return err
	}
	u.Uuid = t.Uuid
	return nil
}

func hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
