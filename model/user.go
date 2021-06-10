package model

import (
	"projectapi/crypt"
	"projectapi/model/entity"
)

type User struct {
	entity.BasicEntity
	name     string
	login    string
	password string
	email    string
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(n string) {
	u.name = n
}

func (u *User) GetLogin() string {
	return u.login
}

func (u *User) SetLogin(l string) {
	u.login = l
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) SetEmail(e string) {
	u.email = e
}

func (u *User) setPassword(p string) {
	u.password = p
}

func (u *User) ChangePassword(d string, e crypt.Encrypter) {
	encrypted := e.EncryptPassword(d)
	u.setPassword(encrypted)
}
