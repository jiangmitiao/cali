package api

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/jiangmitiao/cali/app/services"
	"github.com/revel/revel"
	"strconv"
)

type Tag struct {
	*revel.Controller
}

func (c Tag) Index() revel.Result {
	return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApi())
}

//all tags count
func (c Tag) TagsCount() revel.Result {
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(services.QueryTagsCount()))
}

//all tags info
func (c Tag) Tags() revel.Result {
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(services.QueryTags(limit, start)),
	)
}
