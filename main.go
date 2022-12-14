package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"kietchung/controllers"
	"kietchung/services"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	server       *gin.Engine
	cs           services.ChemistryService
	cc           controllers.ChemistryController
	ctx          context.Context
	chemistryCol *mongo.Collection
	refDocCol    *mongo.Collection
	mongoclient  *mongo.Client
	err          error
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
	mongoConn := fmt.Sprintf("mongodb://%s:%s@localhost:27017/chemistry", username, password)
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

	chemistryCol = mongoclient.Database("chemistry").Collection("chemistry")
	refDocCol = mongoclient.Database("chemistry").Collection("ref_document")

	cs = services.NewUserService(chemistryCol, refDocCol, ctx)
	cc = controllers.New(cs)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	cc.RegisterUserRoutes(basepath)

	server.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		MaxAge: 12 * time.Hour,
	}))
	server.Run(":3000")
}
