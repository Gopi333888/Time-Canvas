package main

import (
	"WatchHive/pkg/config"
	"WatchHive/pkg/di"
	"fmt"
	"log"

	_ "WatchHive/cmd/api/docs"

	"github.com/joho/godotenv"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// @title Go + Gin E-Commerce API Watch Hive
// @version 1.0.0
// @description Watch Hive is an E-commerce platform to purchase Watch
// @contact.name API Support
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the env file")
	}
	config, Err := config.LoadConfig()
	if Err != nil {
		log.Fatal("cannot load config : ", Err)
	}
	fmt.Println("before server")
	server, diErr := di.InitializeAPI(config)
	fmt.Println("After server")
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		fmt.Println("Inside else")
		server.Start()
		fmt.Println("Start")
	}
}
