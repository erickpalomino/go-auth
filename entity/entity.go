package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm: "index", "primaryKey"`
	Name     string
	Lastname string
	Age      int
	Password string
}
