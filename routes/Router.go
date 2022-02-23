package routes

import (
	controller "golang_crudapp/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/api/users", controller.GetUsers())
	incomingRoutes.GET("/api/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/api/users", controller.CreateUser())
	incomingRoutes.PUT("/api/users/:user_id", controller.UpdateUser())
	incomingRoutes.DELETE("/api/users/:user_id", controller.DeleteUser())
}
