package main

import (
	"github.com/gin-gonic/gin"
	"go-auth/controllers/auth"
	"go-auth/database/mongoDB"
	"go-auth/middlewares/authMiddleware"
	"log"
)

func main() {
	err := mongoDB.Connect()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.POST("/auth/login", auth.Login)
	router.POST("/auth/register", auth.Register)
	router.GET("/auth/verifyLogin", authMiddleware.GetAccessToRoute(auth.VerifyLogin))
	log.Fatal(router.Run("localhost:8080"))
}
