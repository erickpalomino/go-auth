package main

import (
	"fmt"
	"service/auth/config"
	"service/auth/routes"
	"service/auth/service"

	"github.com/gin-gonic/gin"
)

func main() {
	err := config.InitGlobalConfiguration("auth-config.yml")
	if err != nil {
		fmt.Println("Error initializing configuration")
	}
	server := gin.Default()
	service.InitDb()
	server.GET("/signup", routes.SignUp)
	server.POST("/encrypt", routes.Encrypt)
	server.POST("/decrypt", routes.Decrypt)
	server.Run()
}
