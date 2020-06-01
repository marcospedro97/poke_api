package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"
)

type Token struct {
	UserUUID  uint32
	Expires   int64
	TokenUUID uint32
}

type Access struct {
	TokenUUID string
	UserUUID  int64
}

func NewToken(u uint32) Token {
	t := Token{UserUUID: u,
		Expires:   time.Now().Add(time.Hour * 12).Unix(),
		TokenUUID: uuid.New().ID()}
	return t
}

func (t *Token) CreateRegister() error {
	u := strconv.FormatUint(uint64(t.UserUUID), 10)
	at := time.Unix(t.Expires, 0)
	tu := strconv.FormatUint(uint64(t.TokenUUID), 10)
	err := Db.R.Set(tu, u, at.Sub(time.Now())).Err()
	if err != nil {
		return err
	}
	return nil
}

func ReadRegister() {
	userUUID := Db.R.Get()
}

func (t *Token) GenerateJWT() (string, error) {
	os.Setenv("API_SECRET", "jdnfksdmfksd")
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = t.UserUUID
	claims["exp"] = t.Expires
	claims["token_id"] = t.TokenUUID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if err := t.CreateRegister(); err != nil {
		return "", err
	}
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
