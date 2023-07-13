package routes

import (
	//"golang-restaurant-backend/controller"

	"golang-restaurant-backend/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(req *gin.Engine){
	req.GET("/user",controller.GetUsers())
	req.GET("/user/:id",controller.GetUser())

	req.POST("/user/signup",controller.SignUp())
	req.GET("/user/login",controller.Login())
}