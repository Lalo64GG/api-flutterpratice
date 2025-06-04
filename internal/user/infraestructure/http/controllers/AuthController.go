package controllers

import (
	middlewares "api/internal/shared/middlewares/jwt"
	"api/internal/shared/response"
	"api/internal/user/application"
	"api/internal/user/infraestructure/http/controllers/helpers"
	"api/internal/user/infraestructure/http/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{
	AuthUseCase *application.AuthUseCase
	BcryptHelper *helpers.BcryptHelper
}

func NewAuthController(authUseCase *application.AuthUseCase) *AuthController {
	return &AuthController{ AuthUseCase: authUseCase}
}

func (ctr *AuthController) Run(ctx *gin.Context){
	var req request.AuthRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Status:  false,
			Message: "Invalid request format",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	user, err := ctr.AuthUseCase.Run(req.Email); 
	
	if err != nil {
		switch err.Error() {
		case "sql: no rows in result set":
			ctx.JSON(http.StatusNotFound, response.Response{
				Status: false,
				Message: "Usuario no encontrado",
				Data:    nil,
				Error:   err.Error(),
			})
			default:
				ctx.JSON(http.StatusInternalServerError, response.Response{
					Status: false,
					Message: "Error al inciar sesión",
					Data:    nil,
					Error:   err.Error(),
				})
			}
			return 
	}

	if err := ctr.BcryptHelper.Compare(user.Password, []byte(req.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, response.Response{
			Status: false,
			Message: "Contraseña incorrecta",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	token, err := middlewares.GenerateJWT(int64(user.ID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Status:  false,
			Message: "Error al generar el token JWT",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}


	ctx.JSON(http.StatusOK, response.Response{
		Status:  true,
		Message: "Inicio de sesión exitoso",
		Data:    map[string]interface{}{
			"token":token,
		},
		Error:   nil,
	})
}