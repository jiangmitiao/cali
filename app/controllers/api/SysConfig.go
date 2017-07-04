package api

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/revel/revel"
)

type SysConfig struct {
	*revel.Controller
}

func (c SysConfig) Index() revel.Result {
	return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApi())
}

func (c SysConfig) Configs() revel.Result {
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(sysConfigService.QuerySysConfigs(1000, 0)),
	)
}

func (c SysConfig) Update() revel.Result {
	id := c.Request.FormValue("id")
	key := c.Request.FormValue("key")
	value := c.Request.FormValue("value")
	sysConfig := models.SysConfig{
		Id:    id,
		Ikey:  key,
		Value: value,
	}
	if ok := sysConfigService.UpdateConfig(sysConfig); ok {
		return c.RenderJSONP(
			c.Request.FormValue("callback"),
			models.NewOKApiWithInfo(sysConfig),
		)
	} else {
		return c.RenderJSONP(
			c.Request.FormValue("callback"),
			models.NewErrorApi(),
		)
	}
}
