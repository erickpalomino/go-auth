package routes

import (
	"net/http"
	"service/auth/entity"
	"service/auth/util/encryption"

	"github.com/gin-gonic/gin"
)

func Encrypt(c *gin.Context) {
	var encryptRequest entity.EncryptRequest
	c.ShouldBindJSON(&encryptRequest)
	encryptedText := encryption.Encrypt(encryptRequest.Text)
	c.JSON(http.StatusOK, gin.H{
		"message": encryptedText,
	})
}

func Decrypt(c *gin.Context) {
	var decryptRequest entity.EncryptRequest
	c.ShouldBindJSON(&decryptRequest)
	decrpytedText := encryption.Decrypt(decryptRequest.Text)
	c.JSON(http.StatusOK, gin.H{
		"message": decrpytedText,
	})
}
