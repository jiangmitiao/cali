package api

import (
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
func (c User) Index(callback string) revel.Result {
	return c.RenderJSONP(callback, models.NewOKApi())
}

func (c User) Login(callback string) revel.Result {
	var loginName string
	var loginPassword string
	c.Params.Bind(&loginName, "loginName")
	c.Params.Bind(&loginPassword, "loginPassword")
	if loginName == "" || loginPassword == "" {
		errStatus := models.NewErrorApiWithMessageAndInfo("用户名或密码错误", nil)
		errStatus.StatusCode = 401
		return c.RenderJSONP(callback, errStatus)
	}

	if userInfo, exist := userService.GetUserByLoginName(loginName); exist && userInfo.LoginPassword == rcali.Sha3_256(loginPassword+userInfo.Salt) {
		//if exist and password correct
		loginSession := rcali.Sha3_256(userInfo.LoginPassword + strconv.FormatInt(time.Now().Unix(), 10))
		userService.FreshLoginSession(loginSession, userInfo.Id)
		return c.RenderJSONP(callback, models.NewOKApiWithMessageAndInfo("login success", loginSession))
	} else {
		errStatus := models.NewErrorApiWithMessageAndInfo("用户名或密码错误", nil)
		errStatus.StatusCode = 402
		return c.RenderJSONP(callback, errStatus)
	}
}

func (c User) Logout(callback, session string) revel.Result {
	rcali.DeleteLoginSession(session)
	return c.RenderJSONP(callback, models.NewOKApi())
}

func (c User) Regist(callback, loginName, loginPassword string) revel.Result {
	return c.RenderJSONP(callback, models.NewOKApi())
}
