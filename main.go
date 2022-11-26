package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"

	"kietchung/controllers"
	"kietchung/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	server      *gin.Engine
	cs          services.ChemistryService
	cc          controllers.ChemistryController
	ctx         context.Context
	userc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

const (
	USERNAME = "username"
	PASSWORD = "password"
)

func init() {
	ctx = context.TODO()
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		return
	}

	username := os.Getenv(USERNAME)
	password := os.Getenv(PASSWORD)
	mongoConn := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.znigccy.mongodb.net/?retryWrites=true&w=majority", username, password)
	mongoconn := options.Client().ApplyURI(mongoConn)
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)

	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userc = mongoclient.Database("chemistry").Collection("chemistry")
	cs = services.NewUserService(userc, ctx)
	cc = controllers.New(cs)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	cc.RegisterUserRoutes(basepath)

	server.Run(":80")
}
