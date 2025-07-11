package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sonu31/experiment-golang-jwt-projct/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"

	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})

	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"sucess": "Access granted for api-2"})

	})

	log.Println("🚀 Server running at http://localhost:" + port)
	router.Run(":" + port)

}

/* {
"First_name":"Sonu",
"Last_name":"Sharma",
"Password":"sonu@1122",
"Email":"sonuhjpkumar@gmail.com",
"Phone":"873317110",
"User_type":"ADMIN"
} */
