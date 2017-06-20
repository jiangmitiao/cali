package api

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/revel/revel"
)

type SysStatus struct {
	*revel.Controller
}

func (c SysStatus) Index() revel.Result {
	return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApi())
}

func (c SysStatus) Status() revel.Result {
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(sysStatusService.QuerySysStatus(1000, 0)),
	)
}
