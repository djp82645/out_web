package user

import (
	// "fmt"
	"errors"

	"api_server/global"
	// "api_server/model/common/request"
	// "api_server/utils"
)

type UserService struct{}

func (exa *UserService) Login(username, password string) (string, error) {
	if userInfo,ok := global.TOKENS[username];ok {
		if userInfo["password"] == password {
			return userInfo["token"], nil
		}
		return "", errors.New("password error")
	}
	return "", errors.New("user not found")

}

func (exa *UserService) SearchUser(token string) (map[string]interface{}, error) {
	if userInfo,ok := global.USERS[token];ok {
		return userInfo, nil
	}
	return nil, errors.New("user not found")
}

func (exa *UserService) Logout() ( error) {
	return nil
}