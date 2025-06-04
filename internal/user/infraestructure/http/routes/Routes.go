package routes

import (
	"api/internal/user/infraestructure/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	createController := http.SetUpCreate()
	getByIdController := http.SetUpGetById()
	authController := http.SetUpAuth()

	router.POST("/", createController.Run)
	router.GET("/:id", getByIdController.Run)
	router.POST("/auth", authController.Run)
}