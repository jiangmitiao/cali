package api

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/services"
	"github.com/revel/revel"
)

type Tag struct {
	*revel.Controller
}

func (c Tag) Index(callback string) revel.Result {
	return c.RenderJSONP(callback, models.NewOKApi())
}

//all tags count
func (c Tag) TagsCount(callback string) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryTagsCount()))
}

//all tags info
func (c Tag) Tags(callback string, limit, start int) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryTags(limit, start)),
	)
}
