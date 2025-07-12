package middleware

import (
	"fmt"
	"net/http"

	helper "github.com/sonu31/experiment-golang-jwt-projct/helpers"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		clinetToken := c.Request.Header.Get("token")
		if clinetToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorizatino header provides")})
			c.Abort()
			return

		}
		claims, err := helper.ValidateToken(clinetToken)

		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err})
			c.Abort()
			return
		}

		// if err != "" {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"err", err})
		// 	c.Abort()
		// 	return

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_Type)
		c.Next()
	}
}
