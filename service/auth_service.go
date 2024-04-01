package service

import (
	"fmt"
	"service/auth/entity"
	"service/auth/util/token"
)

func SignUp(user entity.UserInput) {

	var userDB entity.User
	userDB.Name = user.Name
	userDB.Username = user.Username
	userDB.Age = user.Age
	userDB.Lastname = user.Lastname
	userDB.Password = user.Password
	userDB.HashPassword()
	userDB.SaveThisUser()
}

func SignIn(loginReq entity.LoginRequest) string {
	var userDB entity.User
	entity.GetDatabase().Where("username = ?", loginReq.Username).Find(&userDB)
	if userDB.ComparePassword(loginReq.Password) {
		authToken, err := token.CreateToken(userDB.Username)
		if err != nil {
			fmt.Errorf("token generation failed")
		}
		return authToken
	} else {
		return ""
	}
}
