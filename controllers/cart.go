package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func AddToCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productQueryID := ctx.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")
			_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
		}
	}
}
