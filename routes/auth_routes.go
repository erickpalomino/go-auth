package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "test",
	})
}
