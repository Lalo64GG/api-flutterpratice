package ports

import entites "api/internal/user/domain/entities"

type IUserRepository interface {
	Create(user entites.User) (entites.User, error)
	GetById(id int64) (entites.User ,error)
	GetByEmail(email string) (entites.User, error)
}