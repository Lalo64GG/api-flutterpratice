package helpers

import "golang.org/x/crypto/bcrypt"

type BcryptHelper struct{}

func NewBcryptHelper() (*BcryptHelper, error) {
	return &BcryptHelper{}, nil
}

func (s *BcryptHelper) Encrypt(pwd []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (s *BcryptHelper) Compare(hashPwd string, plainPwd []byte) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPwd), plainPwd)
}