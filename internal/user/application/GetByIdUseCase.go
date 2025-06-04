package application

import (
	"api/internal/user/domain/entities"
	"api/internal/user/domain/ports"
)

type GetByIdUseCase struct {
	UserRepository ports.IUserRepository
}

func NewGetByIdUseCase(userRepository ports.IUserRepository) *GetByIdUseCase{
	return &GetByIdUseCase{
		UserRepository: userRepository,
	}
}

func (u *GetByIdUseCase) Run(id int64) (entities.User, error){
	user, err := u.UserRepository.GetById(id)
	
	if err != nil {
		return entities.User{}, err
	}
	
	return user, nil
}