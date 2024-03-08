package main

import (
	"os"
	  "restaurant/database"
	  "github.com/JUSTICEESSIELP/restuarant/routes"
	  "github.com/JUSTICEESSIELP/restuarant/middlewares"

	  "go.mongodb.org/mongo-driver/mongo"
)



func main(){
port := os.Getenv("PORT")

if port == ""{
	port = "3000"
}



router:=gin.New()




}