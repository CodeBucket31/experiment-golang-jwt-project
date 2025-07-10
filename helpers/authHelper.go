package helpers

import (
	"errors"
	"os/user"

	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUid(c, *gin.cContext ,userId string)(err error){
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "USER" && uid!=  userId{
		err = errors.New("Unauthorized to access this resource")
		return err
	}

err = 	CheckUserType(c, userType)
return err


}