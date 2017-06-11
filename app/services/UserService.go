package services

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"time"
)

type UserService struct {
}

//获取user信息
func (userService UserService) GetUserById(Id string) models.UserInfo {

	//TODO
	return models.UserInfo{Id: Id}

}

func (userService UserService) GetUserByLoginName(loginName string) (models.UserInfo, bool) {
	var userInfo models.UserInfo
	if has, err := localEngine.Where("login_name = ?", loginName).Get(&userInfo); has && err == nil {
		return userInfo, true
	} else {
		return userInfo, false
	}

}

func (userService UserService) FreshLoginSession(loginSession string, UserId string) {
	rcali.SetLoginUser(loginSession, UserId)
}

func (userService UserService) GetLoginUser(loginSession string) (models.UserInfo, bool) {
	id, saveTime := rcali.GetUserIdByLoginSession(loginSession)
	if time.Now().Unix()-saveTime.Unix() > 7200 {
		rcali.DeleteLoginSession(loginSession)
		return models.UserInfo{Id: ""}, false
	} else {
		return userService.GetUserById(id), true
	}
}
