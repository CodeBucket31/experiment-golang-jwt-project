package helper

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sonu31/experiment-golang-jwt-projct/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetailes struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_Type  string
	// jwt.StandardClaims
	jwt.RegisteredClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, usertype string, uid string) (signedToken string, signRefreshToken string) {
	claims := &SignedDetailes{

		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		User_Type:  usertype,
		// StandardClaims: jwt.StandardClaims{
		// 	ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		// },
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	// refreshClains := &SignedDetails{

	// StandardClaims: jwt.StandaredClaims{
	// 	ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(169)).Unix(),
	// },
	// }

	refreshClaims := &SignedDetailes{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(168 * time.Hour)), // 7 days
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken

}

func ValidateToken(signedToken string) (claims *SignedDetailes, msg string) {
	token, err := jwt.ParseWithClaims(signedToken, &SignedDetailes{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil

	})

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetailes)

	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	// if claims.ExpiresAt < time.Now().Local().Unix() {
	// 	msg = fmt.Sprintf("token is expired")
	// 	msg = err.Error()
	// 	return

	// }
	return claims, msg

}

func UpdateAllTokens(signedToken string, signeRefreshToken string, usedId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refresh token", signeRefreshToken})
	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at", Updated_at})

	upsert := true
	filter := bson.M{"user_id": usedId}
	opt := options.UpdateOptions{Upsert: &upsert}

	_, err := userCollection.UpdateOne(ctx, filter, bson.D{
		{"$set", updateObj},
	}, &opt)
	defer cancel()
	if err != nil {
		log.Panic(err)
		return
	}
	return
}
