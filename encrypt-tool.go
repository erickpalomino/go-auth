package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"service/auth/util/encryption"
	"strings"
)

func amain() {
	reader := bufio.NewReader(os.Stdin)

	os := runtime.GOOS

	fmt.Println("Encryption Shell")
	fmt.Println("--------------------------")
	fmt.Println("Write the text to encrypt")
	textToEncrypt, err := reader.ReadString('\n')

	if os == "windows" {
		textToEncrypt = strings.Replace(textToEncrypt, "\r\n", "", -1)
	} else {
		textToEncrypt = strings.Replace(textToEncrypt, "\n", "", -1)
	}

	if err != nil {
		panic(err)
	}
	encryptedText := encryption.Encrypt(textToEncrypt)
	fmt.Printf("Encrypted text: %s \n", encryptedText)

}
