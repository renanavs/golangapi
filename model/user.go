package model

import (
	"projectapi/crypt"
	"projectapi/model/entity"
)

type User struct {
	entity.BasicEntity
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *User) ChangePassword(raw_password string, e crypt.Encrypter) {
	encrypted := e.EncryptPassword(raw_password)
	u.Password = encrypted
}
