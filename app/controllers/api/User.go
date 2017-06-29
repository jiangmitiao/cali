package api

import (
	"github.com/google/uuid"
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/revel/revel"
	"strconv"
	"time"
)

type User struct {
	*revel.Controller
}

// /user
func (c User) Index() revel.Result {
	return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApi())
}

func (c User) Login() revel.Result {
	callback := c.Request.FormValue("callback")
	var loginName string = c.Request.FormValue("loginName")
	var loginPassword string = c.Request.FormValue("loginPassword")
	if loginName == "" || loginPassword == "" {
		errStatus := models.NewErrorApiWithMessageAndInfo(c.Message("loginNameOrLoginPasswordError"), nil)
		errStatus.StatusCode = 401
		return c.RenderJSONP(callback, errStatus)
	}

	if len(loginName) > 64 || len(loginPassword) > 64 {
		errStatus := models.NewErrorApiWithMessageAndInfo(c.Message("loginNameOrLoginPasswordError"), nil)
		errStatus.StatusCode = 402
		return c.RenderJSONP(callback, errStatus)
	}

	if userInfo, exist := userService.GetUserByLoginName(loginName); exist && userInfo.LoginPassword == rcali.Sha3_256(loginPassword+userInfo.Salt) {
		//if exist and password correct
		loginSession := rcali.Sha3_256(userInfo.LoginPassword + strconv.FormatInt(time.Now().Unix(), 10))
		userService.FreshLoginSession(loginSession, userInfo.Id)
		return c.RenderJSONP(callback, models.NewOKApiWithMessageAndInfo(c.Message("loginSuccess"), loginSession))
	} else {
		errStatus := models.NewErrorApiWithMessageAndInfo(c.Message("loginNameOrLoginPasswordError"), nil)
		errStatus.StatusCode = 402
		return c.RenderJSONP(callback, errStatus)
	}
}

//get userinfo by session
func (c User) Info() revel.Result {
	callback := c.Request.FormValue("callback")
	session := c.Request.FormValue("session")
	user, has := userService.GetLoginUser(session)
	if has {
		user.Salt = ""
		user.LoginPassword = ""
		return c.RenderJSONP(callback, models.NewOKApiWithInfo(user))
	} else {
		return c.RenderJSONP(callback, models.NewErrorApiWithMessageAndInfo(c.Message("loginNameOrLoginPasswordError"), nil))
	}

}

//find a session is or not login
func (c User) IsLogin() revel.Result {
	callback := c.Request.FormValue("callback")
	session := c.Request.FormValue("session")
	id, _ := rcali.GetUserIdByLoginSession(session)
	if id == "" {
		return c.RenderJSONP(callback, models.NewErrorApi())
	} else {
		return c.RenderJSONP(callback, models.NewOKApi())
	}
}

//delete the server's login cache
func (c User) Logout() revel.Result {
	callback := c.Request.FormValue("callback")
	session := c.Request.FormValue("session")
	rcali.DeleteLoginSession(session)
	return c.RenderJSONP(callback, models.NewOKApi())
}

//regist a user ,if delete watcherUserRegist in role action ,then not allow to regist
func (c User) Regist() revel.Result {
	callback := c.Request.FormValue("callback")
	loginName := c.Request.FormValue("loginName")
	loginPassword := c.Request.FormValue("loginPassword")

	if !c.Validation.Check(loginName, revel.ValidEmail()).Ok {
		return c.RenderJSONP(callback, models.NewErrorApiWithMessageAndInfo(c.Message("signupfail")+" - "+c.Message("needemail"), nil))
	}

	if loginName == "" || loginPassword == "" || len(loginName) > 64 || len(loginPassword) > 64 {
		return c.RenderJSONP(callback, models.NewErrorApiWithMessageAndInfo(c.Message("signupfail")+c.Message("loginNameOrLoginPasswordError"), nil))
	} else {
		salt := uuid.New().String()
		safePassword := rcali.Sha3_256(loginPassword + salt)
		newUser := models.UserInfo{
			Id:            uuid.New().String(),
			LoginName:     loginName,
			LoginPassword: safePassword,
			Salt:          salt,
			UserName:      loginName,
			Email:         loginName,
		}
		returnmessage := c.Message("signupsuccess")
		if rcali.HasNeedActive() {
			newUser.Valid = 2
			returnmessage = c.Message("pleaseactive")
		}
		if userService.Regist(newUser) {
			rcali.SendActiveMail(newUser.Email, newUser.Salt)
			return c.RenderJSONP(callback, models.NewOKApiWithMessageAndInfo(returnmessage, nil))
		} else {
			return c.RenderJSONP(callback, models.NewErrorApiWithMessageAndInfo(c.Message("signupfail")+c.Message("loginNameOrLoginPasswordError"), nil))
		}
	}
}

// update userName and email by this method
func (c User) Update() revel.Result {
	callback := c.Request.FormValue("callback")
	session := c.Request.FormValue("session")
	userName := c.Request.FormValue("userName")
	if len(userName) > 64 {
		return c.RenderJSONP(callback, models.NewErrorApiWithMessageAndInfo(c.Message("emailOrUsernameIsTooLong"), nil))
	}
	if user, isLogin := userService.GetLoginUser(session); isLogin {
		user.UserName = userName
		user.Img = ""
		if updateOK := userService.UpdateInfo(user); updateOK {
			return c.RenderJSONP(callback, models.NewOKApi())
		} else {
			return c.RenderJSONP(callback, models.NewErrorApiWithMessageAndInfo(c.Message("uncatchedError"), nil))
		}
	} else {
		return c.RenderJSONP(callback, models.NewErrorApiWithMessageAndInfo(c.Message("needLogin"), nil))
	}
}

// change the password ,need oldpassword and newpassword
func (c User) ChangePassword() revel.Result {
	callback := c.Request.FormValue("callback")
	session := c.Request.FormValue("session")
	oldLoginPassword := c.Request.FormValue("oldLoginPassword")
	loginPassword := c.Request.FormValue("loginPassword")
	if user, isLogin := userService.GetLoginUser(session); isLogin {
		if user.LoginPassword == rcali.Sha3_256(oldLoginPassword+user.Salt) {
			//oldpassword is ok
			user.Salt = uuid.New().String()
			user.LoginPassword = rcali.Sha3_256(loginPassword + user.Salt)
			if changed := userService.UpdatePassword(user); changed {
				return c.RenderJSONP(callback, models.NewOKApi())
			} else {
				return c.RenderJSONP(callback, models.NewErrorApiWithMessageAndInfo("uncatched error", nil))
			}
		} else {
			return c.RenderJSONP(callback, models.NewErrorApiWithMessageAndInfo("old password error", nil))
		}
	} else {
		return c.RenderJSONP(callback, models.NewErrorApiWithMessageAndInfo("no login", nil))
	}
}

func (c User) QueryUserCount() revel.Result {
	callback := c.Request.FormValue("callback")
	session := c.Request.FormValue("session")
	if user, isLogin := userService.GetLoginUser(session); isLogin {
		role := userRoleService.GetRoleByUser(user.Id)
		if role.Name == "admin" {
			return c.RenderJSONP(callback, models.NewOKApiWithInfo(userService.QueryUserCount(c.Request.FormValue("loginName"))))
		}
	}
	return c.RenderJSONP(callback, models.NewErrorApi())
}

func (c User) QueryUser() revel.Result {
	callback := c.Request.FormValue("callback")
	session := c.Request.FormValue("session")
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.UserListNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))

	if user, isLogin := userService.GetLoginUser(session); isLogin {
		role := userRoleService.GetRoleByUser(user.Id)
		if role.Name == "admin" {
			return c.RenderJSONP(callback, models.NewOKApiWithInfo(userService.QueryUser(c.Request.FormValue("loginName"), limit, start)))
		}
	}
	return c.RenderJSONP(callback, models.NewErrorApi())
}

func (c User) Delete() revel.Result {
	callback := c.Request.FormValue("callback")
	session := c.Request.FormValue("session")
	userId := c.Request.FormValue("userId")

	if user, isLogin := userService.GetLoginUser(session); isLogin {
		role := userRoleService.GetRoleByUser(user.Id)
		if role.Name != "admin" {
			//delete login user
			go rcali.DeleteLoginUserId(userId)
			return c.RenderJSONP(callback, models.NewOKApiWithInfo(userService.DeleteUser(userId)))
		}
	}
	return c.RenderJSONP(callback, models.NewErrorApi())
}

func (c User) UserStatus() revel.Result {
	user, _ := userService.GetLoginUser(c.Request.FormValue("session"))
	if user.Id != "" {
		tmp := time.Now()
		start := time.Date(tmp.Year(), tmp.Month(), tmp.Day(), 0, 0, 0, 0, tmp.Location())
		stop := start.AddDate(0, 0, 1)
		config, _ := userConfigService.GetUserConfig(user.Id)
		status := make(map[string]string)
		status["count"] = strconv.Itoa(userService.GetDownloadCount(user.Id, start, stop))
		status["maxcount"] = strconv.Itoa(config.MaxDownload)
		return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApiWithInfo(status))
	} else {
		return c.RenderJSONP(c.Request.FormValue("callback"), models.NewErrorApi())
	}

}

func (c User) Active() revel.Result {
	key := c.Request.FormValue("key")
	if userService.ActiveUser(key) {
		return c.Redirect("/login")
	}
	return c.RenderJSONP(c.Request.FormValue("callback"), models.NewErrorApi())
}
