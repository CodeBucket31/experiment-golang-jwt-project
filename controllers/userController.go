package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator.v10"
	"github.com/go-playground/validator/v10"
	"github.com/sonu31/experiment-golang-jwt-projct/database"
	"github.com/sonu31/experiment-golang-jwt-projct/helpers"
	"github.com/sonu31/experiment-golang-jwt-projct/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string) string{
	bcrypt.GenerateFromPassword([]byte(password),14)
	if err != nil {
		log.Panic(err)
		
	}
	return  string(bytes)
}




func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("email of password is incorrect")
		check = false

	}
	return check, msg

}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validator.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.ConntDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {

			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while"})

		}

		password := HashPassword(*user.Passwoed)
		user.Passwoed = &password

		count, err := userCollection.ConntDocuments(ctx, bson.M{"phone": user.Phone})

		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured whilte checkeing for "})
			return nil, err
		}

		if count >= 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this emil or phone number "})
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format((time.RFC3339)))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ := helper.GenerateAllTokens(*&user.Email, *&user.First_name, *&user.Last_name, *&user.User_Type, *&user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)

		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)

	}

}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Passwoed, *foundUser.Passwoed)
		defer cancel()

		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError,gin.H{"error",msg})
			return 
			
		}

		if foundUser.Email == nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"user not found"})
		}


    token ,refreshToken , _ :=  helper.GenerateAllTokens(*foundUser.Email ,*foundUser.First_name,*&foundUser.Last_name,*foundUser.User_Type,foundUser.User_id)
     helper.UpdateAllTokens(token ,refreshToken,foundUser.User_id)
	err =  userCollection.FindOne(ctx,bson.M{"user_id":foundUser.User_id}).Decode(&foundUser)
	 if err != nil {
        c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error})
       return 
	}
	c.JSON(http.StatusOK,foundUser)


	}
}

func GetUsers()gin.HandlerFunc{
	return  func(ctx *gin.Context) {
		helpers.CheckUserType(c,"ADMIN"); err != nil{
        c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
		}
		var ctx , cancel = context.WithTimeout(context.Background(),time.Second)
	recordPerPage ,err = 	strconv.Atoi(c.Query("recordPerPage"))
   if err != nil || recrecordPerPage <1 {
	recrecordPerPage = 10

     }
	 strconv.Atoi(c.Query("page")) 

	}

}



func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
			return
		}
		c.JSON(http.StatusOK, user)

	}
}
