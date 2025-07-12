package routes

import (
	controller "github.com/sonu31/experiment-golang-jwt-projct/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("user/login", controller.Login())
}
