package api

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/services"
	"github.com/revel/revel"
)

type Author struct {
	*revel.Controller
}

func (c Author) Index(callback string) revel.Result {
	return c.RenderJSONP(callback, models.NewOKApi())
}

//all tags count
func (c Author) AuthorsCount(callback string) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryAuthorsCount()))
}

//all tags info
func (c Author) Authors(callback string, limit, start int) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryAuthors(limit, start)),
	)
}
