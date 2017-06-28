package services

import (
	"github.com/google/uuid"
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"time"
)

type UserService struct {
}

//获取user信息
func (userService UserService) GetUserById(Id string) models.UserInfo {
	var userInfo = models.UserInfo{}
	engine.ID(Id).Where("valid = ?", 0).Get(&userInfo)
	return userInfo
}

func (userService UserService) GetUserByLoginName(loginName string) (models.UserInfo, bool) {
	var userInfo = models.UserInfo{}
	if has, err := engine.Where("login_name = ?", loginName).Where("valid = ?", 0).Get(&userInfo); has && err == nil {
		return userInfo, true
	} else {
		return userInfo, false
	}
}

func (userService UserService) GetAllUserByLoginName(loginName string) (models.UserInfo, bool) {
	var userInfo = models.UserInfo{}
	if has, err := engine.Where("login_name = ?", loginName).Get(&userInfo); has && err == nil {
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
	_, has := userService.GetAllUserByLoginName(user.UserName)
	if !has {
		session := engine.NewSession()
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

func (userService UserService) QueryUserCount(name string) int64 {
	var count int64 = 0
	if name != "" {
		count, _ = engine.Cols("id", "login_name", "user_name", "email", "img").Where("login_name like ?", "%"+name+"%").Or("user_name like ?", "%"+name+"%").Where("valid = ?", 0).Count(&models.UserInfo{})
	} else {
		count, _ = engine.Cols("id", "login_name", "user_name", "email", "img").Where("valid = ?", 0).Count(&models.UserInfo{})
	}
	return count
}

//find user by username or login name
func (userService UserService) QueryUser(name string, limit, start int) []models.UserInfo {
	users := make([]models.UserInfo, 0)
	if name != "" {
		engine.Cols("id", "login_name", "user_name", "email", "img").Where("login_name like ?", "%"+name+"%").Or("user_name like ?", "%"+name+"%").Where("valid = ?", 0).Limit(limit, start).Find(&users)
	} else {
		engine.Cols("id", "login_name", "user_name", "email", "img").Where("valid = ?", 0).Limit(limit, start).Find(&users)
	}
	return users
}

//set valid = 1 ,not allow delete admin
func (userService UserService) DeleteUser(userId string) bool {
	user := userService.GetUserById(userId)
	if user.Id != "" && user.LoginName != "admin" {
		user.Valid = 1
		_, err := engine.Id(userId).Cols("valid").Update(user)
		if err != nil {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

//update info .not in password or other.
func (userService UserService) UpdateInfo(user models.UserInfo) bool {
	_, err := engine.ID(user.Id).Cols("user_name", "img").Where("valid = ?", 0).Update(user)
	if err != nil {
		return false
	} else {
		return true
	}
}

//change password and salt
func (userService UserService) UpdatePassword(user models.UserInfo) bool {
	_, err := engine.ID(user.Id).Cols("login_password", "salt").Where("valid = ?", 0).Update(user)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (userService UserService) AddUpload(userId string, bookId int) bool {
	upload := models.UserInfoBookUploadLink{Id: uuid.New().String(), UserInfo: userId, Book: bookId, CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}
	if _, err := engine.InsertOne(upload); err == nil {
		return true
	} else {
		return false
	}
}

func (userService UserService) AddDownload(userId string, bookId string) bool {
	download := models.UserInfoBookDownloadLink{Id: uuid.New().String(), UserInfo: userId, Book: bookId, CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}
	if _, err := engine.InsertOne(download); err == nil {
		return true
	} else {
		return false
	}
}

func (userService UserService) GetDownloadCount(userId string, start, stop time.Time) int {
	count, _ := engine.Where("user_info = ?", userId).And("created >= ?", start.Unix()).And("created <= ?", stop.Unix()).Count(models.UserInfoBookDownloadLink{})
	return int(count)
}

func (userService UserService) ActiveUser(salt string) bool {
	userinfo := models.UserInfo{}
	engine.Where("salt = ?", salt).Get(&userinfo)
	if userinfo.Salt == salt {
		userinfo.Valid = 0
		engine.Where("salt = ?", salt).Cols("valid").Update(userinfo)
		return true
	} else {
		return false
	}
}
