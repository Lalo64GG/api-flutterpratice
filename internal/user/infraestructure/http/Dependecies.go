package http

import (
	"api/internal/user/application"
	"api/internal/user/application/services"
	"api/internal/user/domain/ports"
	"api/internal/user/infraestructure/adapters"
	"api/internal/user/infraestructure/http/controllers"
	"api/internal/user/infraestructure/http/controllers/helpers"
	"log"
)

var (
	userRepository ports.IUserRepository
	bcryptService services.Bcrypt
)

func init() {
	var err error
	  userRepository, err = adapters.NewUserRepositoryMySql()
    if err != nil {
        log.Fatalf("Error initializing user repository: %v", err)
    }

    // Inicializar el servicio de encriptación
    bcryptService, err = helpers.NewBcryptHelper()
    if err != nil {
        log.Fatalf("Error initializing bcrypt service: %v", err)
    }
}


func SetUpCreate() *controllers.CreateController{
	createUseCase := application.NewCreateUseCase(userRepository, bcryptService)
	return controllers.NewCreateController(createUseCase)
}

func SetUpAuth() *controllers.AuthController {
	authUseCase := application.NewAuthUseCase(userRepository)
	return controllers.NewAuthController(authUseCase)
}

func SetUpGetById() *controllers.GetByIdController {
	getByIdUseCase := application.NewGetByIdUseCase(userRepository)
	return controllers.NewGetByIdController(getByIdUseCase)
}