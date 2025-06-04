package controllers

import (
	"api/internal/shared/response"
	"api/internal/user/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetByIdController struct {
	GetByIdUseCase *application.GetByIdUseCase
}

func NewGetByIdController(getByIdUseCase *application.GetByIdUseCase) *GetByIdController {
	return &GetByIdController{ GetByIdUseCase: getByIdUseCase }
}

func (ctr *GetByIdController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Status:  false,
			Message: "Invalid ID format",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	user, err := ctr.GetByIdUseCase.Run(id)

	if err != nil{ 
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Status:  false,
			Message: "Error retrieving user",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Status:  true,
		Message: "User retrieved successfully",
		Data:    user,
		Error:   nil,
	})
}