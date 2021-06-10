package crypt

type Encrypter interface {
	EncryptPassword(p string) string
}

type BCrypt struct {
}

func (b *BCrypt) EncryptPassword(p string) string {
	return ""
}
