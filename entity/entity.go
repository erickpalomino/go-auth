package entity

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Name     string
	Lastname string
	Age      int
	Password string
}

func (u *User) HashPassword() string {
	passwordBytes := []byte(u.Password)
	hashedPassword := sha256.Sum256(passwordBytes)
	u.Password = fmt.Sprintf("%X", hashedPassword)
	return u.Password
}

func (u *User) SaveThisUser() {
	GetDatabase().Save(&u)
}

func (u *User) ComparePassword(password string) bool {
	passwordBytes := []byte(password)
	hashedPassword := sha256.Sum256(passwordBytes)
	password = fmt.Sprintf("%X", hashedPassword)
	return strings.Compare(u.Password, password) == 0
}
