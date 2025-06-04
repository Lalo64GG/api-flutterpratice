package application

import (
	"api/internal/user/domain/entities"
	"api/internal/user/domain/ports"
)

type AuthUseCase struct {
	UserRepository ports.IUserRepository
}

func NewAuthUseCase(userRepository ports.IUserRepository) *AuthUseCase{
	return &AuthUseCase{UserRepository: userRepository}
}

func (u *AuthUseCase) Run(email string) (entities.User, error){
	user, err := u.UserRepository.GetByEmail(email)

	if err != nil {
		return entities.User{}, err
	}

	if user.ID == 0 {
		return entities.User{}, nil
	}

	return user, nil
}