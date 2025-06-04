package services

type Bcrypt interface {
	Encrypt(pwd []byte) (string, error)
	Compare(hashedPwd string, plainPwd []byte)  error
}