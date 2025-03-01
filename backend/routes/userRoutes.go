package routes

import (
	"crud_usuarios/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
}
