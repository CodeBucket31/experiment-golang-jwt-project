package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sonu31/experiment-golang-jwt-projct/routes"
)

func main() {
	port = os.Getenv("PORT")

	if port == "" {
		port = "8000"

	}

	router := gin.New()
	router.User(gin.Logger())

	routes.AuthRouter(router)
	routes.UserRouter(router)

	router.GET("api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})

	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"sucess": "Access granted for api-2"})

	})

}
