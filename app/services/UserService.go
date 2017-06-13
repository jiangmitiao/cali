package services

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
)

type UserService struct {
}

//获取user信息
func (userService UserService) GetUserById(Id string) models.UserInfo {
	var userInfo = models.UserInfo{}
	localEngine.ID(Id).Where("valid = ?",0).Get(&userInfo)
	return userInfo
}

func (userService UserService) GetUserByLoginName(loginName string) (models.UserInfo, bool) {
	var userInfo = models.UserInfo{}
	if has, err := localEngine.Where("login_name = ?", loginName).Where("valid = ?",0).Get(&userInfo); has && err == nil {
		return userInfo, true
	} else {
		return userInfo, false
	}

}

func (userService UserService) FreshLoginSession(loginSession string, UserId string) {
	rcali.SetLoginUser(loginSession, UserId)
}

func (userService UserService) GetLoginUser(loginSession string) (models.UserInfo, bool) {
	id, _ := rcali.GetUserIdByLoginSession(loginSession)
	if id == "" {
		return models.UserInfo{}, false
	} else {
		return userService.GetUserById(id), true
	}
}

func (userService UserService) Regist(user models.UserInfo) bool {
	_, has := userService.GetUserByLoginName(user.UserName)
	if !has {
		if _, err := localEngine.Insert(user); err == nil {
			return true
		}
	}
	return false
}
