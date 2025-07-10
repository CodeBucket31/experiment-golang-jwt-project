package routes


import (
"github.com/gin-gonic/gin"
controller "experiment-golang-jwt-projct/controllers"
)

func AuthRoutes(incoming Routes  *gin.Engine){
	incomingRoutes.POST("users/signup",controller.Signup())
	incomingRoutes.POST("user/login",controller.Login())
}