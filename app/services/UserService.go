package services

import (
	"github.com/google/uuid"
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
)

type UserService struct {
}

//获取user信息
func (userService UserService) GetUserById(Id string) models.UserInfo {
	var userInfo = models.UserInfo{}
	localEngine.ID(Id).Where("valid = ?", 0).Get(&userInfo)
	return userInfo
}

func (userService UserService) GetUserByLoginName(loginName string) (models.UserInfo, bool) {
	var userInfo = models.UserInfo{}
	if has, err := localEngine.Where("login_name = ?", loginName).Where("valid = ?", 0).Get(&userInfo); has && err == nil {
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
		session := localEngine.NewSession()
		defer session.Close()
		// add Begin() before any action
		err := session.Begin()

		user.Id = uuid.New().String()
		if _, err := session.Insert(user); err == nil {

		} else {
			session.Rollback()
			return false
		}

		userRole := models.UserInfoRoleLink{
			Id:       uuid.New().String(),
			UserInfo: user.Id,
			Role:     "user",
		}

		if _, err := session.Insert(userRole); err == nil {

		} else {
			session.Rollback()
			return false
		}

		// add Commit() after all actions
		err = session.Commit()
		if err != nil {
			panic(err)
		}
		return true
	}
	return false
}

//update info .not in password or other.
func (userService UserService) UpdateInfo(user models.UserInfo) bool {
	_, err := localEngine.ID(user.Id).Cols("user_name", "email", "img").Where("valid = ?", 0).Update(user)
	if err != nil {
		return false
	} else {
		return true
	}
}

//change password and salt
func (userService UserService) UpdatePassword(user models.UserInfo) bool {
	_, err := localEngine.ID(user.Id).Cols("login_password", "salt").Where("valid = ?", 0).Update(user)
	if err != nil {
		return false
	} else {
		return true
	}
}
