package main

import (
	"os"

	"restaurant/database"
	"restaurant/routes"

	"github.com/gin-gonic/gin"

	middlewares "restaurant/middlewares"

	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	ginRouter := gin.New()
	ginRouter.Use(gin.Logger())
	ginRouter.Use(middlewares.Authentication())

	routes.FoodRoutes(ginRouter)
	routes.MenuRoutes(ginRouter)
	routes.InvoiceRoutes(ginRouter)
	routes.TableRoutes(ginRouter)
	routes.OrderItemRoutes(ginRouter)
	routes.OrderRoutes(ginRouter)
	routes.UserRoutes(ginRouter)

	ginRouter.Run(":" + port)

}
