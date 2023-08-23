package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marioheryanto/erajaya/go-app/controller"
	"github.com/marioheryanto/erajaya/go-app/database"
	"github.com/marioheryanto/erajaya/go-app/helper"
	"github.com/marioheryanto/erajaya/go-app/library"
	"github.com/marioheryanto/erajaya/go-app/repository"
	"github.com/marioheryanto/erajaya/go-app/route"
)

func init() {
	godotenv.Load()
}

func main() {
	// clients
	dbClient := database.ConnectDB()
	redisClient := database.ConnectRedis()
	validator := helper.NewValidator()

	// repo
	productRepo := repository.NewProductRepository(dbClient, redisClient)

	// library
	productLib := library.NewProductLibrary(productRepo, validator)

	// controller
	productCtrl := controller.NewProductController(productLib)

	router := gin.Default()

	// config
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	// config.AllowHeaders = []string{}

	// ----- Router -----
	router.Use(cors.New(config))

	route.ProductRoutes(router, productCtrl)

	// router.Run(fmt.Sprintf("http://localhost:%v", os.Getenv("PORT")))
	router.Run(":" + os.Getenv("PORT"))

}
