package main

import (
	"fmt"
	"service/auth/config"
	"service/auth/entity"
	"service/auth/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	err := config.InitGlobalConfiguration("auth-config.yml")
	if err != nil {
		fmt.Println("Error initializing configuration")
	}
	server := gin.Default()
	entity.InitDb()
	entity.GetDatabase()
	server.POST("/signup", routes.SignUp)
	server.POST("/signin", routes.SignIn)
	server.POST("/encrypt", routes.Encrypt)
	server.POST("/decrypt", routes.Decrypt)
	server.Run()
}
