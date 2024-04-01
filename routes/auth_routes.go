package routes

import (
	"net/http"
	"service/auth/entity"
	"service/auth/service"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user entity.UserInput
	c.ShouldBindJSON(&user)
	service.SignUp(user)
	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

func SignIn(c *gin.Context) {
	var loginReq entity.LoginRequest
	c.ShouldBindJSON(&loginReq)
	token := service.SignIn(loginReq)
	if len(token) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"token":   token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "login failed",
		})
	}

}
