package main

import (
	"golang-restaurant-backend/routes"
	"log"
	"os"

	//"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	//"golang-restaurant-backend/middleware"
	"golang-restaurant-backend/db"
)



func main() {

    client, err := db.DbConnect()
     if err != nil {
		log.Fatal(err)
		return
	}

     port:=os.Getenv("PORT")

	 if(port!=""){
		port="8080"
	 }

	 router := gin.New()

  	 router.Use(gin.Logger())

	 
    //router.Use(middleware.Auth())
//use Authentication Middleware.Note that  You cant use  below Routes without Authentication
//Routes Come Here From Packages.

    routes.UserRoutes(router)


	router.Run(":"+port)
}