package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"
)

var (
	//needs to be % 16 = 0
	secretKey string = "0123456789123456"
	iv        string = "0000000000000000"
)

func Encrypt(text string) string {
	if len(text)%16 != 0 {
		padding := 16 - (len(text) % 16)
		text = text + strings.Repeat(" ", padding)
	}

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	cipher := cipher.NewCBCEncrypter(block, []byte(iv))
	var encryptedTextBytes = make([]byte, len(text))
	cipher.CryptBlocks(encryptedTextBytes, []byte(text))

	return string(base64.StdEncoding.EncodeToString(encryptedTextBytes))
}

func Decrypt(encryptedText string) string {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	cipher := cipher.NewCBCDecrypter(block, []byte(iv))
	decodedEncryptedBytes, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		panic(err)
	}
	var decryptedTextBytes = make([]byte, len(decodedEncryptedBytes))
	cipher.CryptBlocks(decryptedTextBytes, decodedEncryptedBytes)
	decryptedText := string(decryptedTextBytes[:])
	decryptedText = strings.ReplaceAll(decryptedText, " ", "")
	return decryptedText
}
