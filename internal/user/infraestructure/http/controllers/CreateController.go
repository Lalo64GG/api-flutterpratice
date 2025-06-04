package controllers

import (
	"api/internal/shared/response"
	"api/internal/user/application"
	"api/internal/user/infraestructure/http/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateController struct {
	CreateUseCase *application.CreateUseCase
	Validator *validator.Validate
}

func NewCreateController(createUseCase *application.CreateUseCase) *CreateController{
	return &CreateController{
		CreateUseCase: createUseCase,
		Validator: validator.New(),
	}
}

func (ctr  *CreateController ) Run(ctx *gin.Context){
	var req request.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Status:  false,
			Message: "Invalid request format",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := ctr.Validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Status:  false,
			Message: "Validation error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	user, err := ctr.CreateUseCase.Run(req.Name, req.LastName, req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Status:  false,
			Message: "Error creating user",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Status:  true,
		Message: "User created successfully",
		Data:    user,
		Error:   nil,
	})

}

