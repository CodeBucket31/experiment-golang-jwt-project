package helper

import (
	"errors"
	"go/token"
	"os"
	"time"

	"github.com/sonu31/experiment-golang-jwt-projct/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SignedDetailes struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_Type  string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client,"user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string,firstName string, lastName string, usertype string,uid string)(signedToken string,signRefreshToken string){
	claims := &SignedDetailes{

		Email :email,
		First_name : firstName,
		Last_name :lastName,
		Uid:uid,
		User_Type:userusertype,
		StandardClaims :jwt.StandardClaims{
			ExpiresAt:time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
		

	}
	refreshClains  := &SignedDetails{

		StandardClaims :jwt.StandaredClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(169)).Unix(),
		},
	}

 token ,err :=  jwt.NewWithClaims.jwt.SigningMethodHS256,claims).SignedString([]byte(SECRET_KEY))
 refreshToken ,err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]bype(SECRET_Key))


 if err != nil {
  log.Panic(err)
  return
 }

 return token , refreshToken,err



}

