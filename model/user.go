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

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(n string) {
	u.Name = n
}

func (u *User) GetLogin() string {
	return u.Login
}

func (u *User) SetLogin(l string) {
	u.Login = l
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetEmail(e string) {
	u.Email = e
}

func (u *User) setPassword(p string) {
	u.Password = p
}

func (u *User) ChangePassword(raw_password string, e crypt.Encrypter) {
	encrypted := e.EncryptPassword(raw_password)
	u.setPassword(encrypted)
}
