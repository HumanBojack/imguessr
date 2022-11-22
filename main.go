package main

import (
	"context"
	"imguessr/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World!")
}

func main() {
	router := gin.Default()

	router.GET("/hello", getHello)

	router.GET("/second", getHello)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("imguessr")
	usersCollection := db.Collection("users")

	models.Seed(usersCollection, ctx)


	router.Run("0.0.0.0:8080")
}
