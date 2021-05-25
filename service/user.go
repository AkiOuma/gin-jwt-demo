package service

import (
	"fmt"
)

type User struct {
	Userid   string `json:"userid"`
	Password string `json:"-"`
	Group    int    `json:"group"`
}

var UsersMook = []User{
	{"yuki", "123", 1},
	{"kira", "456", 0},
	{"shinn", "789", 2},
}

func GetUser(userid string, password string) (user User, err error) {
	for _, v := range UsersMook {
		if v.Userid == userid && v.Password == password {
			user = v
			return
		}
	}
	err = fmt.Errorf("error: user [%v] not found", userid)
	return
}

func GetAllUser() []User {
	return UsersMook
}
