package application

import (
	"api/internal/user/application/services"
	"api/internal/user/domain/entities"
	"api/internal/user/domain/ports"
)

type CreateUseCase struct {
	UserRepository ports.IUserRepository
	BycryptService services.Bcrypt
}

func NewCreateUseCase(userRepository ports.IUserRepository, bycryptService services.Bcrypt) *CreateUseCase{
	return &CreateUseCase{
		UserRepository: userRepository,
		BycryptService: bycryptService,
	}
}

func (u *CreateUseCase) Run(Name, LastName, Email, Password string) ( entities.User, error){
	
	hashedPassword, err := u.BycryptService.Encrypt([]byte(Password))

	if err != nil {
		return entities.User{}, err
	}


	user := entities.User{
		Name:     Name,
		LastName: LastName,
		Email:    Email,
		Password: hashedPassword,
	}

	user, err = u.UserRepository.Create(user)

	if err != nil {
		return  entities.User{}, err
	}

	return user, nil
}