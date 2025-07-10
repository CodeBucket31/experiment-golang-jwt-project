package routes

import (
	controller "experiment-golang-jwt-projct/controllers"

	"github.com/gin-gonic/gin"
	"github.com/sonu31/experiment-golang-jwt-projct/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUsers())
}
